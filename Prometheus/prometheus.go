package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Prometheus/http"
	"github.com/luckykris/Cronus/Hephaestus/os"
)

func main() {
	mainCfg := cfg.LoadCfg()
	if mainCfg.Daemon{
		os.Daemonize()
	}
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Init(mainCfg)
	log.Info("Start ", cfg.SOFTWARE, " success")
	http.Start()
	select {
	}
}


