package database

import (
	"database/sql"
	"fmt"
	"github.com/belson77/Go-001/Week04/news/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(cf config.Config) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cf.DB.Username, cf.DB.Password, cf.DB.Host, cf.DB.Dataname))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
