package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Prometheus/http"
	"github.com/luckykris/Cronus/Hephaestus/go/os"
	"github.com/luckykris/Cronus/Hephaestus/go/safe"
)

func main() {
	mainCfg := cfg.LoadCfg()
	if mainCfg.Daemon{
		os.Daemonize()
	}
	log.SetLevel(mainCfg.LogCfg.LevelId)
	//start safe-exiter
	ex:=safe.NewExit()
	prometheus.Init(mainCfg)
	log.Debug("Init ","prometheus "  ," success")
	//
	//ex.Join()
	go http.Start(mainCfg.Debug)
	//
	log.Debug("Start ","http "  ," success")
	log.Info("Start ", cfg.SOFTWARE, " success")
	//set exiter callback for os`s signal handle
	os.StartSignalHandle("interrupt",func(){ex.StartAllExit()},1)
	log.Info("Stoping ", cfg.SOFTWARE, " success")
	ex.WaitAllExit()
	log.Info("Stopped ", cfg.SOFTWARE, " success")
}


