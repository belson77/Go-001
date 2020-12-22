package service

import (
	"context"
	"fmt"
	pb "github.com/belson77/Go-001/Week04/news/api/comment/appcomment/v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewAppCommentServer(conn *grpc.ClientConn) *AppCommentServer {
	return &AppCommentServer{cli: pb.NewAppCommentClient(conn)}
}

type AppCommentServer struct {
	pb.UnimplementedAppCommentServer
	cli pb.AppCommentClient
}

func (srv *AppCommentServer) Submit(ctx context.Context, req *pb.SubmitRequest) (resp *pb.SubmitResponse, err error) {
	fmt.Println("http service: submit")

	resp, err = srv.cli.Submit(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "app comment client request Submit error")
	}
	return
}

func (srv *AppCommentServer) Query(ctx context.Context, req *pb.QueryRequest) (resp *pb.QueryResponse, err error) {
	fmt.Println("http service: query")

	resp, err = srv.cli.Query(ctx, req)
	if err != nil {
		err = errors.Wrap(err, "app comment client requesy Query error")
	}
	return resp, nil
}
