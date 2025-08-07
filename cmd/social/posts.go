package main

import (
	"net/http"
	"strconv"

	"github.com/Suikyoo/go-social/internal/jsonutils"
	"github.com/Suikyoo/go-social/internal/repository"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
  UserID int64 `json:"-"`
  Title string `json:"title"`
  Content string `json:"content"`
}

func (payload *CreatePostPayload) Scan(value authContextValue, w http.ResponseWriter, r *http.Request) error{

  return nil
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {


  val, ok := r.Context().Value(authContextKey).(authContextValue)

  if !ok {
    http.Error(w, "Auth Context value not found", http.StatusInternalServerError)
    return
  }

  //fill in data for postpayload item
  payload := CreatePostPayload{}

  err := jsonutils.Read(w, r, &payload)

  if err != nil {
    http.Error(w, "Invalid Payload", http.StatusNotAcceptable)
    return

  }

  payload.UserID = val.UserID


  //create post item
  post := repository.Post{
    UserID: payload.UserID,
    Title: payload.Title,
    Content: payload.Content,
  }

  //send it to the repository
  err = app.store.Posts.Create(r.Context(), &post)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)

}

func (app *application) getPost(w http.ResponseWriter, r *http.Request) {
  idKey := chi.URLParam(r, "id")
  id, err := strconv.Atoi(idKey)
  if err != nil {
    http.Error(w, "User not found. Invalid ID", http.StatusBadRequest)
    return
  }

  post, err := app.store.Posts.Get(r.Context(), int64(id))

  if err != nil {
    http.Error(w, "Database error", http.StatusInternalServerError)
    return

  }
  jsonutils.Write(w, post)

}

func (app *application) getPostFeed(w http.ResponseWriter, r *http.Request) {
  var feedAmt int8 = 20
  feed, err := app.store.Posts.GetFeed(r.Context(), feedAmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jsonutils.Write(w, feed)
}
