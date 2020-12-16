package data

type CommentRepo struct {
	Id       int64
	UserName string
	Content  string
}

func NewCommentData() *CommentRepo {
	return &CommentData{}
}

func (rp *CommentRepo) Add() (int64, error) {

}
