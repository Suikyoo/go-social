package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestCreatePost() {
  
  body := `{"title": "Test", "content": "this is my test"}`
  req := httptest.NewRequest(http.MethodPost, "/posts", strings.NewReader(body))
  req.Header.Set("Content-Type", "application/json")

  //work in progress (maybe not?)
}
