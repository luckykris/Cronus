package db

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
)

type iDb interface {
	Start() error
	Ping() error
	//Get()
	//Set()
}

func Init(cfg cfg.DbCfgStruct) (iDb, error) {
	switch cfg.Class {
	case "mysql":
		return &MysqlDb{Host: cfg.Host, Port: cfg.Port, User: cfg.User, Passwd: cfg.Passwd, Db: cfg.Db, Charset: cfg.Charset, MaxLifeTime: cfg.MaxLifeTime, MaxIdleConns: cfg.MaxIdleConns}, nil
	default:
		return nil, fmt.Errorf("Don`t support DB class %s ", cfg.Class)
	}
}
