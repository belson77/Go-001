package service

import (
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
