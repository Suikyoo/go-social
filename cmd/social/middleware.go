package main

import (
	"net/http"
	"strings"
)



func AuthTokenMiddleware (next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
      //send failed response
    }

    parts := strings.Split(authHeader, " ")
    if len(parts) !=2 || parts[0] != "Bearer" {
      //send failed response
    }

    //do some verification with the parts[1] aka the jwt string token
    //wip
  })
}

