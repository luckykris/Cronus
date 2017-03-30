package db

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"strings"
	//"github.com/luckykris/Cronus/Prometheus/global"
)


type CRUD interface{
	Get(string, interface{},[]string, []string, ...interface{}) (*Cur, error)
	GetJoin(string, [][4]string, interface {}, []string, []string, ...interface {}) (*Cur, error)
	//GetCount(string, interface{} , []string)(int,error)
	Add(string, []string, [][]interface{}) error
	Delete(string, []string) error
	Update(string, []string, []string, []interface{}) error
}

type Dbi interface {
	Start() error
	Ping() error
	Begin()(Txi,error)
	CRUD
}

type Txi interface {
	Commit() error
	Rollback() error
	CRUD
}
const (
	LEFT string="LEFT"
	RIGHT string="RIGHT"
	INNER string="INNER"
)




func Init(cfg cfg.DbCfgStruct) (Dbi, error) {
	switch strings.ToLower(cfg.Class) {
	case "mysql":
		return &MysqlDb{Host: cfg.Host, Port: cfg.Port, User: cfg.User, Passwd: cfg.Passwd, Db: cfg.Db, Charset: cfg.Charset, MaxLifeTime: cfg.MaxLifeTime, MaxIdleConns: cfg.MaxIdleConns}, nil
	default:
		return nil, fmt.Errorf("Don`t support DB class %s ", cfg.Class)
	}
}
