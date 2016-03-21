package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	_ "github.com/luckykris/go-utilbox/exit"
	"os"
)

func main() {
	mainCfg := cfg.LoadCfg()
	log.SetLevel(mainCfg.LogCfg.LevelId)
	log.Debug("Start init Database.")
	dbobj, err := db.Init(mainCfg.DbCfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
	log.Debug("Success init Database.")
	log.Debug("Start Open Database.")
	err = dbobj.Start()
	if err != nil {
		log.Error("Open Database failed=>", err.Error())
		os.Exit(255)
	}
	log.Debug("Open Database Success")
	log.Info("Start ", cfg.SOFTWARE, " success")
	for {
	}
}
