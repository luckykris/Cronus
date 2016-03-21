package db

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDb struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	Db           string
	Charset      string
	MaxLifeTime  int64
	MaxIdleConns int64
	DbPool       *sql.DB
}

func (db *MysqlDb) Start() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", db.User, db.Passwd, db.Host, int(db.Port), db.Db, db.Charset)
	log.Debug("DSN==>", dsn)
	DbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DbPool = DbPool
	//dr := time.ParseDuration(fmt.Sprintf("%ds", t))
	db.DbPool.SetConnMaxLifetime(time.Duration(db.MaxLifeTime) * time.Second)
	db.DbPool.SetMaxIdleConns(int(db.MaxIdleConns))
	return db.Ping()
}
func (db *MysqlDb) Ping() error {
	return db.DbPool.Ping()
}
