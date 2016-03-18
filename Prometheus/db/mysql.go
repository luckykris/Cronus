package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	User    string
	Passwd  string
	Db      string
	Charset string
	Conn		mysql.driver.Conn
}

func (db *MysqlDb) Start() {
	Conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=%s",db.User,db.Passwd,db.Db,db.Charset)	
	if err != nil {
		return
	}	
	db.Conn = Conn	
}


