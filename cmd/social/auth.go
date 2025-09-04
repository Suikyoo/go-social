package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Suikyoo/go-social/internal/authutils"
	"github.com/Suikyoo/go-social/internal/jsonutils"
	"github.com/Suikyoo/go-social/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//changed my mind I should probably make something like an independent auth internal package huhu

//ok, made an authutils package just to
//decouple the password type along with its methods


type ReceivedUserPayload struct {
	Username string   `json:"username"`
	Password authutils.Password `json:"password"`
}

type SentTokenPayload struct {
	Token string `json:"token"`
	Type  string `json:"type"`
	//in seconds
	ExpiresIn int64 `json:"expires_in"`
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	payload := ReceivedUserPayload{}
	err := jsonutils.Read(w, r, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = app.store.Users.GetByUsername(r.Context(), payload.Username)
	//meaning, there is a result, (an account with the same username)
	if err == nil {
		http.Error(w, "Username already taken", http.StatusConflict)
		return
	}

	user := repository.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	err = app.store.Users.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return 
	}

	w.WriteHeader(http.StatusCreated)

  _, err = w.Write([]byte("User successfully created"))
  if err != nil {
    log.Println("Response not sent", err.Error())
  }

}

func (app *application) createToken(w http.ResponseWriter, r *http.Request) {
	userPayload := ReceivedUserPayload{}

	if err := jsonutils.Read(w, r, &userPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := app.store.Users.GetByUsername(r.Context(), userPayload.Username)
	//meaning, there is a result, (an account with the same username)
	if err != nil {
		http.Error(w, "No such user", http.StatusConflict)
    return
	}

	//if user is there,

  //if database user's password, is the same as the sent user's password,
  if !user.Password.Compare(&userPayload.Password){
    http.Error(w, "Wrong password", http.StatusUnauthorized)
    return
  }

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatInt(user.ID, 10),
		"exp": time.Now().Add(time.Second * time.Duration(app.config.auth.tokenExpiry)).Unix(),
	})

	jwtString, err := token.SignedString(app.config.auth.jwtSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return 

	}

  //use cookies instead of a json response
  /*
	tokenPayload := SentTokenPayload{
		Token:     jwtString,
		Type:      "Bearer",
		ExpiresIn: app.config.auth.tokenExpiry,
	}

	jsonutils.Write(w, tokenPayload)
	//send the token back
  */

  cookie := http.Cookie{
    Name: "social-auth-token",
    Value: "Bearer " + jwtString,
    Path: "/",

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
		Expires: time.Now().Add(24 * time.Hour),
    MaxAge: 86400,
    Secure: true,
    HttpOnly: true,
    SameSite: http.SameSiteLaxMode,

  }
  http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusCreated)

  _, err = w.Write([]byte("Token successfully created"))
  if err != nil {
    log.Println("Response not sent", err.Error())
  }

}
