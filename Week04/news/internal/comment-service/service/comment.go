package service

import (
	"context"
	"database/sql"
	pb "github.com/belson77/Go-001/Week04/news/api/comment/appcomment/v1"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/biz"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/data"
)

func NewCommentService(db *sql.DB) *CommentService {
	return &CommentService{dao: db}
}

type CommentService struct {
	pb.UnimplementedAppCommentServer
	dao *sql.DB
}

func (s *CommentService) Submit(ctx context.Context, req *pb.SubmitRequest) (*pb.SubmitResponse, error) {

	// repo init
	repo := data.NewCommentRepo(s.dao)
	cu := biz.NewCommentUsecase(repo)

	commentDo := &biz.Comment{
		ObjID:    req.GetObjID(),
		ObjType:  int(req.GetObjType()),
		Content:  req.GetContent(),
		UserName: req.GetUserName(),
	}

	id, err := cu.SubmitComment(commentDo)
	if err != nil {
		return nil, err
	}
	return &pb.SubmitResponse{ID: id}, nil
}

func (s *CommentService) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {

	// repo init
	repo := data.NewCommentRepo(s.dao)
	cu := biz.NewCommentUsecase(repo)

	comments, err := cu.QueryComment(req.GetObjID(), int(req.GetObjType()))
	if err != nil {
		return nil, err
	}

	resp := &pb.QueryResponse{}
	for _, v := range comments {
		item := &pb.QueryResponse_QueryItem{}
		item.ID = v.ID
		item.ObjID = v.ObjID
		item.ObjType = int32(v.ObjType)
		item.UserName = v.UserName
		item.Content = v.Content
		resp.Items = append(resp.Items, item)
	}
	return resp, nil
}
