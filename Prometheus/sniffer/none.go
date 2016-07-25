package sniffer

import (
	"time"
	//"runtime"
	//log "github.com/Sirupsen/logrus"
	//"github.com/luckykris/Cronus/Hephaestus/safe"
)

type None struct{
	Interval time.Duration
}

//func (sniffer *None)Start(*safe.Exiter){
//	exit.Join()
//	for {
//		runtime.Gosched() 
//		select {
//		case <-time.After(sniffer.Interval * time.Minute):
//			log.Debug("start sniffer.")
//		}
//		log.Debug("finish sniffer.")
//	}
//}

func (sniffer *None)Run() error{
	return nil
}
func (sniffer *None)GetInterval()time.Duration{
	return sniffer.Interval
}