package database

import (
	"database/sql"
	"fmt"
	"github.com/belson77/Go-001/Week04/news/config"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string
	Username string
	Password string
	Dataname string
}

func NewMysql(cf config.Config) (*sql.DB, error) {
	dbCf, _ := cf.(MysqlConfig)
	conn, err := sql.Open("mysql", fmt.Sprint("%s:%s@%s/%s", dbCf.Username, dbCf.Password, dbCf.Host, dbCf.Dataname))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
