package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	_ "github.com/luckykris/go-utilbox/exit"
	"os"
)

func main() {
	mainCfg := cfg.LoadCfg()
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Start()
	log.Info("Start ", cfg.SOFTWARE, " success")
	for {
	}
}
