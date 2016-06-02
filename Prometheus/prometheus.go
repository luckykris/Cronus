package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Prometheus/http"
	_ "github.com/luckykris/go-utilbox/exit"
)

func main() {
	mainCfg := cfg.LoadCfg()
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Init(mainCfg)
	log.Info("Start ", cfg.SOFTWARE, " success")
	http.Start()
	select {
	}
}
