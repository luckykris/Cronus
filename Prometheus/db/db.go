package db

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"strings"
	//"github.com/luckykris/Cronus/Prometheus/global"
)

type Dbi interface {
	Start() error
	Ping() error
	Get(string, []string, []string, ...interface{}) (*Cur, error)
	Add(string, []string, [][]interface{}) error
	Delete(string, []string) error
	Update(string, []string, []string, []interface{}) error
}

func Init(cfg cfg.DbCfgStruct) (Dbi, error) {
	switch strings.ToLower(cfg.Class) {
	case "mysql":
		return &MysqlDb{Host: cfg.Host, Port: cfg.Port, User: cfg.User, Passwd: cfg.Passwd, Db: cfg.Db, Charset: cfg.Charset, MaxLifeTime: cfg.MaxLifeTime, MaxIdleConns: cfg.MaxIdleConns}, nil
	default:
		return nil, fmt.Errorf("Don`t support DB class %s ", cfg.Class)
	}
}
