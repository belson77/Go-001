// +build wireinject

package main

import (
	"context"
	"github.com/belson77/Go-001/Week04/news/config/yaml"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/service"
	app "github.com/belson77/Go-001/Week04/news/internal/comment/service"
	"github.com/belson77/Go-001/Week04/news/internal/pkg/database"
	"github.com/belson77/Go-001/Week04/news/internal/pkg/grpc"
	"github.com/google/wire"
)

func initializeCommentService(ctx context.Context, f string) (*service.CommentService, error) {
	wire.Build(service.NewCommentService, database.NewMysql, yaml.NewConfig)
	return nil, nil
}

func initializeCommentApp(ctx context.Context, f string) (*app.AppCommentServer, error) {
	wire.Build(app.NewAppCommentServer, grpc.NewClient, yaml.NewConfig)
	return nil, nil
}
