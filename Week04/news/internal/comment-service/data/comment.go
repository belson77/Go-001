package data

import (
	"database/sql"
	"github.com/belson77/Go-001/Week04/news/internal/biz"
	_ "github.com/go-sql-driver/mysql"
)

var _ biz.CommentRepo = (commentRepo)(nil)

func NewCommentRepo() *comentRepo {
	conn, err := sql.Open("mysql", "root:123456@/comment")
	if err != nil {
		return nil
	}
	return &commentRepo{conn}
}

type commentRepo struct {
	db *sql.DB
}

func (cm *commentRepo) Add(o *biz.Comment) error {
	sql := "INSERT INTO `comment`(obj_id,obj_type,user_id,content) VALUE(?,?,?,?)"
	if err := cm.db.Query(sql); err != nil {
		err = errors.Wrap(err, sql)
		return err
	}
	return nil
}
