package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Prometheus/http"
	"github.com/luckykris/Cronus/Prometheus/sniffer"
	"github.com/luckykris/Cronus/Hephaestus/os"
	"github.com/luckykris/Cronus/Hephaestus/safe"
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
	go http.Start()
	log.Debug("Start ","http "  ," success")
	//start sniffer
	ex.Join()
	go sniffer.Start(mainCfg.SnifferCfg,ex)
	log.Debug("Start ","sniffer "  ," success")
	log.Info("Start ", cfg.SOFTWARE, " success")
	//set exiter callback for os`s signal handle
	os.StartSignalHandle("interrupt",func(){ex.StartAllExit()},3)
	log.Info("Stoping ", cfg.SOFTWARE, " success")
	ex.WaitAllExit()
	log.Info("Stopped ", cfg.SOFTWARE, " success")
}


