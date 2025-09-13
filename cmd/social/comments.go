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
		RequestError(w, "Invalid Payload")
		return
	}

	postID, err := strconv.Atoi(r.PathValue("postID")) 
	if err != nil {
		InternalError(w, err)
		return
	}

	ctxValue, err := authContextKey.Get(r.Context())
	if err != nil {
		InternalError(w, err)
		return
	}
	comment := repository.Comment{
		Content: payload.Content, 
		UserID: ctxValue.UserID, 
		PostID: int64(postID),
	}
	//repository function create comment
	err = app.store.Comments.Create(r.Context(), &comment)
	if err != nil {
		DBError(w)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (app *application) getCommentFeed(w http.ResponseWriter, r *http.Request) {

	postID, err := strconv.ParseInt(r.PathValue("postID"), 10, 0)
	if err != nil {
		InternalError(w, err)
		return
	}

	feedAmt := int8(20)

  feed, err := app.store.Comments.GetFeedByPostID(r.Context(), postID, feedAmt)
  if err != nil {
		DBError(w)
    return
  }

  jsonutils.Write(w, feed)
}

