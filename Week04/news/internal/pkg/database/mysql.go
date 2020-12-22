package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/belson77/Go-001/Week04/news/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(ctx context.Context, cf config.Config) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cf.DB.Username, cf.DB.Password, cf.DB.Host, cf.DB.Dataname))
	if err != nil {
		return nil, err
	}
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("MySQL Close")
			conn.Close()
		}
	}()
	return conn, nil
}
