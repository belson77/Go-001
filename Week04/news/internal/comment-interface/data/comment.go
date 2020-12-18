package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	DB *sql.DB
}

func NewData() (*data, error) {
	conn, err := sql.Open("mysql", "root:123456@/comment")
	if err != nil {
		return nil, err
	}
	return &data{conn}, nil
}
