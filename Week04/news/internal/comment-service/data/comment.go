package data

import (
	"database/sql"
	"github.com/belson77/Go-001/Week04/news/internal/comment-service/biz"
	"github.com/pkg/errors"
)

var _ biz.CommentRepo = (*commentRepo)(nil)

func NewCommentRepo(db *sql.DB) biz.CommentRepo {
	return &commentRepo{db}
}

type commentRepo struct {
	db *sql.DB
}

func (cm *commentRepo) Add(o *biz.Comment) (id int64, err error) {
	sql := "INSERT INTO `comment`(obj_id,obj_type,user_name,content) VALUE(?,?,?,?)"
	res, err := cm.db.Exec(sql, o.ObjID, o.ObjType, o.UserName, o.Content)
	if err != nil {
		err = errors.Wrap(err, sql)
		return
	}

	id, err = res.LastInsertId()
	return
}
