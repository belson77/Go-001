package biz

import (
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/data"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/service"
)

// DO
type Comment struct {
	Id       int64
	ObjId    int64
	ObjType  int64
	UserName string
	Content  string
}

type CommentRepo interface {
	Submit(*Comment)
}

func NewCommentUsecase(repo *CommentRepo) *CommentUsecase {
	return &CommentUsecase
}

type CommentUsecase struct {
	repo *CommentRepo
}

func (c comm) SubmitComment(d *CommentRepo) error {
	comm := d.(data.CommentRepo)
	id, err := comm.Add()
	if err != nil {
		return err
	}
	return nil
}
