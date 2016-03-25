package db

import (
	"bytes"
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luckykris/Cronus/Prometheus/global"
	//"reflect"
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
	TABLEdeviceType string = `device_type`
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

func (db *MysqlDb) GetDeviceType(args ...string) (interface{}, error) {
	sql := fmt.Sprintf(`SELECT device_model_id ,device_model_name FROM %s`, TABLEdeviceType)
	allDeviceType := []global.DeviceType{}
	rows, err := db.DbPool.Query(sql)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return allDeviceType, err
		}
		allDeviceType = append(allDeviceType, global.DeviceType{Id: id, Name: name})
	}
	return allDeviceType, nil
}

func (db *MysqlDb) Find(table string, columns_name [][]byte, columns ...interface{}) (*Cur, error) {
	sql := fmt.Sprintf(`SELECT %s FROM %s`, bytes.Join(columns_name, []byte(`,`)), TABLEdeviceType)
	rows, err := db.DbPool.Query(sql)
	return &Cur{rows, columns_name, columns}, err

}
func (c *Cur) Fetch() bool {
	if c.row.Next() {
		err := c.row.Scan(c.columns...)
		if err != nil {
			return false
		}
		return true
	} else {
		c.row.Close()
		return false
	}
}
