package prometheus

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	"os"
)

type Prometheus struct {
	dbobj  db.Dbi
}

var PROMETHEUS *Prometheus

func Init(mainCfg cfg.MainCfg) {
	var err error
	log.Debug("Start init Database.")
	PROMETHEUS = &Prometheus{}
	PROMETHEUS.dbobj, err = db.Init(mainCfg.DbCfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
	log.Debug("Success init Database.")
	log.Debug("Start Open Database.")
	err = PROMETHEUS.dbobj.Start()
	if err != nil {
		log.Error("Open Database failed=>", err.Error())
		os.Exit(255)
	}
	log.Debug("Open Database Success")
}

func Start() {

}
