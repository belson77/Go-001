package biz

import (
	"news/internal/comment/data"
	"news/internal/comment/service"
)

type CommentRepo interface {
	Add()
}

func AddComment(d *CommentRepo) error {
	comm := d.(data.CommentRepo)
	id, err := comm.Add()
	if err != nil {
		return err
	}
	return nil
}
