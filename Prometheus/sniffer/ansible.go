package sniffer

import (
	"time"
	"runtime"
	"os/exec"
	log "github.com/Sirupsen/logrus"
)

type Ansible struct{
	Interval time.Duration
	Exe  string
}

func (sniffer *Ansible)Start(){
	for {
		runtime.Gosched() 
		select {
		case <-time.After(sniffer.Interval * time.Minute):
			log.Debug("start sniffer.")
			err:=sniffer.Run()
			if err!=nil{
				log.Warn("command run failed:",err.Error())
			}
		}
		log.Debug("finish sniffer.")
	}
}


func (sniffer *Ansible)Run()error{
	c:=exec.Command(sniffer.Exe,"192.168.33.81,192.168.33.82")
	r,err:=c.Output()
	log.Debug(string(r))
	return err
}