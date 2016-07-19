package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Prometheus/http"
	"github.com/luckykris/Cronus/Prometheus/sniffer"
	"github.com/luckykris/Cronus/Hephaestus/os"
)

func main() {
	mainCfg := cfg.LoadCfg()
	if mainCfg.Daemon{
		os.Daemonize()
	}
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Init(mainCfg)
	log.Debug("Init ","prometheus "  ," success")
	go http.Start()
	log.Debug("Start ","http "  ," success")
	sn,err:=sniffer.Init(mainCfg.SnifferCfg)
	if err!=nil{
		log.Error("Init ","sniffer "  ," failed:", err.Error())
		return
	}else{
		log.Debug("Init ","sniffer "  ," success")
	}
	go sn.Start()
	log.Info("Start ", cfg.SOFTWARE, " success")
	select {
	}
}


