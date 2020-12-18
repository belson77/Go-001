package service

import (
	"database/sql"
	"fmt"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/biz"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/data"
	"net/http"
	"strconv"
)

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{db}
}

type CommentService struct {
	dao *sql.DB
}

func (svc *CommentService) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PostFormValue("obj_id"), 10, 64)
	objType, _ := strconv.Atoi(r.PostFormValue("obj_type"))

	// repo init
	repo := data.NewCommentRepo(svc.dao)
	cu := biz.NewCommentUsecase(repo)

	commentDo := &biz.Comment{
		0,
		id,
		objType,
		r.PostFormValue("user_name"),
		r.PostFormValue("content"),
	}

	id, err := cu.SubmitComment(commentDo)
	if err != nil {
		w.Write([]byte("提交失败"))
	}
	w.Write([]byte(fmt.Sprintf("提交成功. id=%d", id)))
}
