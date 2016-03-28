package db

import (
	"bytes"
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
type Cur struct {
	row          *sql.Rows
	columns_name [][]byte
	columns      []interface{}
}

const (
	TABLEdeviceModel string = `device_model`
)

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

func (db *MysqlDb) Close() error {
	return db.DbPool.Close()
}

func (db *MysqlDb) Find(table string, columns_name [][]byte, conditions [][]byte, columns ...interface{}) (*Cur, error) {
	var conditions_str = ""
	if len(conditions) != 0 {
		conditions_str = " WHERE " + string(bytes.Join(conditions, []byte(` AND `)))
	}
	sql := fmt.Sprintf(`SELECT %s FROM %s %s`, bytes.Join(columns_name, []byte(`,`)), TABLEdeviceModel, conditions_str)
	rows, err := db.DbPool.Query(sql)
	return &Cur{rows, columns_name, columns}, err

}
func (c *Cur) Fetch() bool {
	if c.row.Next() {
		err := c.row.Scan(c.columns...)
		if err != nil {
			log.Fatal("db fetch failed:", err.Error())
			return false
		}
		return true
	} else {
		c.row.Close()
		return false
	}
}
