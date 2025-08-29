package main

import (
	"net/http"
	"strconv"

	"github.com/Suikyoo/go-social/internal/jsonutils"
	"github.com/go-chi/chi/v5"
)

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
  idKey := chi.URLParam(r, "id")
  id, err := strconv.Atoi(idKey)

  if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
  }

  user, err := app.store.Users.Get(r.Context(), int64(id))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  jsonutils.Write(w, user)

}

func (app *application) getUserFeed(w http.ResponseWriter, r *http.Request) {
  var feedAmt int8 = 20
  feed, err := app.store.Users.GetFeed(r.Context(), feedAmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jsonutils.Write(w, feed)

}
