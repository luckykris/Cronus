package sniffer
import (
	"strings"
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Hephaestus/safe"
	log "github.com/Sirupsen/logrus"
	"time"
	"path/filepath"
	"runtime"
)


const (
	AnsiblePlugin string = "ansible-sniff.py"
)

type Snifferi interface{

	//Start(*safe.Exiter)
	Run()error
	GetInterval()time.Duration
}

func Init(cfg cfg.SnifferCfgStruct) (Snifferi, error) {
	switch strings.ToLower(cfg.Class) {
	case "none":
		return &None{Interval:time.Duration(cfg.Interval)},nil
	case "ansible":
		return &Ansible{Interval:time.Duration(cfg.Interval),Exe:filepath.Join(cfg.PluginPath,AnsiblePlugin)},nil
	default:
		return nil, fmt.Errorf("unknow sniffer: %q", cfg.Class)
	}
}

func Start(cfg cfg.SnifferCfgStruct,exit *safe.Exiter){
	sniffer,err:=Init(cfg)
	if err!=nil{
		log.Errorf("Init sniffer failed:%s",err.Error())
		exit.FinishOneExit()
		return
	}
	for {
		runtime.Gosched() 
		select {
		case <-exit.WaitExitSignal():
			log.Debug("stoping sniffer")
			exit.FinishOneExit()
			log.Debug("stopped sniffer")
			return
		case <-time.After(sniffer.GetInterval() * time.Minute):
			log.Debug("start sniffer.")
			err:=sniffer.Run()
			if err!=nil{
				log.Warn("command run failed:",err.Error())
			}
		}
		log.Debug("finish sniffer.")
	}
}