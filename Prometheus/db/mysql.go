package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	Host    string
	User    string
	Passwd  string
	Db      string
	Charset string
	Conn    *sql.DB
}

func (db *MysqlDb) Start() error {
	Conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=%s", db.User, db.Passwd, db.Host, db.Db, db.Charset))
	if err != nil {
		return err
	}
	db.Conn = Conn
	return nil
}
