package sniffer

import (
	"time"
	"fmt"
	"runtime"
	"os/exec"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Hephaestus/simplejson"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"github.com/luckykris/Cronus/Hephaestus/net"
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
		serial:=tmp_json.Get("ansible_facts").Get("ansible_product_serial").MustString()
		hostname:=tmp_json.Get("ansible_facts").Get("ansible_hostname").MustString()
		os:=tmp_json.Get("ansible_facts").Get("ansible_distribution").MustString()
		memsize:=tmp_json.Get("ansible_facts").Get("ansible_memtotal_mb").MustInt()
		release:=tmp_json.Get("ansible_facts").Get("ansible_distribution_version").MustFloat64()
		interface_names:=tmp_json.Get("ansible_facts").Get("ansible_interfaces").MustStringArray()
		netPorts:=[]prometheus.NetPort{}
		for _,interface_name :=range interface_names {
			if interface_name == "lo"{
				continue
			}
			mac:=tmp_json.Get("ansible_facts").Get(interface_name).Get("macaddress").MustString()
			_type:=tmp_json.Get("ansible_facts").Get(interface_name).Get("type").MustString()
			ipv4_str:=tmp_json.Get("ansible_facts").Get(fmt.Sprintf("ansible_%s",interface_name)).Get("ipv4").Get("address").MustString()
			ipv4,err:=net.Ipv4StringConverUint32(ipv4_str)
			if err!=nil{
				log.Warn("sniffer convert ipv4 failed",err.Error())
				continue
			}
			netPorts=append(netPorts,prometheus.NetPort{Mac:mac,Ipv4Int:ipv4,Type:_type })
			server:=prometheus.Server{Serial:serial,Hostname:hostname,Os:os,Release:release,Memsize:memsize,NetPorts:netPorts}
		}
		servers=append(servers,server)
	}
	return servers,nil
}