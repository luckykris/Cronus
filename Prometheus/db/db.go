package db

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	//"github.com/luckykris/Cronus/Prometheus/global"
)

type Dbi interface {
	Start() error
	Ping() error
	GetDeviceType(args ...string) (interface{}, error)
	Find(string, [][]byte, ...interface{}) (*Cur, error)
	//Set()
}

func Init(cfg cfg.DbCfgStruct) (Dbi, error) {
	switch cfg.Class {
	case "mysql":
		return &MysqlDb{Host: cfg.Host, Port: cfg.Port, User: cfg.User, Passwd: cfg.Passwd, Db: cfg.Db, Charset: cfg.Charset, MaxLifeTime: cfg.MaxLifeTime, MaxIdleConns: cfg.MaxIdleConns}, nil
	default:
		return nil, fmt.Errorf("Don`t support DB class %s ", cfg.Class)
	}
}
