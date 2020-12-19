// +build wireinject

package main

import (
	"github.com/belson77/Go-001/Week04/news/config/yaml"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/service"
	"github.com/belson77/Go-001/Week04/news/internal/pkg/database"
	"github.com/google/wire"
)

func initializeCommentService(f string) (*service.CommentService, error) {
	wire.Build(service.NewCommentService, database.NewMysql, yaml.NewConfig)
	return nil, nil
}
