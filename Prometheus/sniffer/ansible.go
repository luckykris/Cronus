package sniffer

import (
	"time"
	"fmt"
	"runtime"
	"os/exec"
	log "github.com/Sirupsen/logrus"
	"github.com/bitly/go-simplejson"
	"github.com/luckykris/Cronus/Hephaestus/simplejson"
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
	r,err:=RunAnsibleCommand(sniffer.Exe,"192.168.33.81,192.168.33.82")
	if err!=nil{
		return err
	}
	_,err=DecodeAnsibleReturn(r)
	if err!=nil {
		return err
	}
	return nil
}

func RunAnsibleCommand(exe , arg string)([]byte,error){
	c:=exec.Command(exe,arg)
	r,err:=c.Output()
	return r,err
}

func DecodeAnsibleReturn(ar []byte)([]prometheus.Server,error){
	servers :=[]prometheus.Server{}
	json, err := simplejson.NewJson(ar)
	success_servers,err:=json.Get("contacted").Map()
	if err!=nil{
		return servers,err
	}
	for _,infomation :=range success_servers{
		tmp_json:=&simplejson.Json{Data: infomation,}
		fmt.Println(tmp_json)
		//tmp_json.data=infomation
		serial:=tmp_json.Get("ansible_facts").Get("ansible_product_serial").MustString()
		hostname:=tmp_json.Get("ansible_facts").Get("ansible_hostname").MustString()
		os:=tmp_json.Get("ansible_facts").Get("ansible_os_family").MustString()
		servers=append(servers,prometheus.Server{Serial:serial,Hostname:hostname,Os:os})
	}
	return servers,nil
}