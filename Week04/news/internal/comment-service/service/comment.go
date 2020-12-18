package service

import (
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/biz"
	"net/http"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	commentDo := &biz.Comment{
		r.PostFormValue("obj_id"),
		r.PostFormValue("ojb_type"),
		r.PostFormValue("user_name"),
		r.PostFormValue("content"),
	}
	if err := biz.CommentRepo(commentDo); err != nil {
		w.Write([]byte("提交失败"))
	}
	w.Write([]byte("提交成功"))
}
