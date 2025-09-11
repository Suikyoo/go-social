package main

import (
	"net/http"
	"strconv"

	"github.com/Suikyoo/go-social/internal/jsonutils"
	"github.com/Suikyoo/go-social/internal/repository"
)

type createCommentPayload struct {
	Content string `json:"content"`;
}
func (app *application) createComment(w http.ResponseWriter, r *http.Request) {
	payload := createCommentPayload{}
	err := jsonutils.Read(w, r, &payload)
	if err != nil {
		http.Error(w, "Invalid Payload", http.StatusNotAcceptable)
		return
	}

	postID, err := strconv.Atoi(r.PathValue("postID")) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	ctxValue, err := authContextKey.Get(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	comment := repository.Comment{Content: payload.Content, UserID: ctxValue.UserID, PostID: int64(postID)}
	//repository function create comment
	app.store.Comments.Create(r.Context(), &comment)

}

func (app *application) getCommentFeed(w http.ResponseWriter, r *http.Request) {
  var feedAmt int8 = 20
  feed, err := app.store.Comments.GetFeed(r.Context(), feedAmt)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  jsonutils.Write(w, feed)
}

