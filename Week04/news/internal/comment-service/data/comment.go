package data

import (
	"database/sql"
	"fmt"
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

func (cm *commentRepo) Add(o *biz.Comment) (int64, error) {
	sql := "INSERT INTO `comment`(obj_id,obj_type,user_name,content) VALUE(?,?,?,?)"
	res, err := cm.db.Exec(sql, o.ObjID, o.ObjType, o.UserName, o.Content)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("comment repo sql: %s", sql))
	}

	return res.LastInsertId()
}

func (cm *commentRepo) GetAll(objID int64, objType int) ([]*biz.Comment, error) {
	sql := "SELECT id,obj_id,obj_type,username,content FROM `comment` WHERE `obj_id`=? AND `obj_type`=?"
	rows, err := cm.db.Query(sql, objID, objType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("comment repo sql: %s", sql))
	}
	defer rows.Close()

	res := make([]*biz.Comment, 0)
	for rows.Next() {
		v := biz.Comment{}
		if err = rows.Scan(&v.ID, &v.ObjID, &v.ObjType, &v.UserName, &v.Content); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("comment repo sql: %s", sql))
		}
		res = append(res, &v)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("comment repo sql: %s", sql))
	}
	return res, nil
}
