package sniffer

import (
	"time"
	"runtime"
	log "github.com/Sirupsen/logrus"
)

type None struct{
	Interval time.Duration
}

func (sniffer *None)Start(){
	for {
		runtime.Gosched() 
		select {
		case <-time.After(sniffer.Interval * time.Minute):
			log.Debug("start sniffer.")
		}
		log.Debug("finish sniffer.")
	}
}