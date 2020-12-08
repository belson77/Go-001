package dao

import (
	"fmt"
	"time"
	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_SQL = "SELECT id,name,age FROM user WHERE id=?"
)

func New() *Dao {
	conn, err := sql.Open("mysql", "root:123456@tcp(localhost)/test")
	if err != nil {
		panic(err)
	}
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	return &Dao{conn}
}

type Dao struct {
	db *sql.DB
}

type User struct {
	Id int64 `json:id`
	Name string `json:name`
	Age int `json:age`
}


func (d *Dao) GetUser(uid int64) (*User, error) {
	u := &User{}
	err := d.db.QueryRow(USER_SQL, uid).Scan(&u.Id, &u.Name, &u.Age)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		err = errors.Wrap(err, fmt.Sprintf("SQL: %s; Param: %d", USER_SQL, uid))
		return nil, err
	default:
		return u, nil
	}
}

func (d *Dao) Close() {
	d.db.Close()
}

