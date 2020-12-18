package biz

// DO
type Comment struct {
	ID       int64
	ObjID    int64
	ObjType  int
	UserName string
	Content  string
}

type CommentRepo interface {
	Add(*Comment) (int64, error)
}

func NewCommentUsecase(repo CommentRepo) *CommentUsecase {
	return &CommentUsecase{repo}
}

type CommentUsecase struct {
	repo CommentRepo
}

func (cu *CommentUsecase) SubmitComment(c *Comment) (int64, error) {
	return cu.repo.Add(c)
}
