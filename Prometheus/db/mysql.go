package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	Host    string
	Port    int64
	User    string
	Passwd  string
	Db      string
	Charset string
	Conn    *sql.DB
}

func (db *MysqlDb) Start() error {
	Conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s?charset=%s", db.User, db.Passwd, db.Host, int(db.Port), db.Db, db.Charset))
	if err != nil {
		return err
	}
	db.Conn = Conn
	return nil
}
func (db *MysqlDb) Ping() error {
	return db.Conn.Ping()
}
