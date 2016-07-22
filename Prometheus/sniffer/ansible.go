package sniffer

import (
	"time"
	"fmt"
	"strings"
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
		//select {
		//case <-time.After(sniffer.Interval * time.Minute):
		//	log.Debug("start sniffer.")
			err:=sniffer.Run()
			if err!=nil{
				log.Warn("command run failed:",err.Error())
			}
		//}
		log.Debug("finish sniffer.")
	}
}


func (sniffer *Ansible)Run()error{
	servers,err:=GetServerArray()
	if err!=nil{
		return err
	}
	err=sniffer.SniiferServer(servers)
	if err!=nil {
		return err
	}
	return nil
}

func GetServerArray()([]*prometheus.Server,error){
	return prometheus.GetServer()
}

func RunAnsibleCommand(exe , arg string)([]byte,error){
	log.Debug("sniffer run command:",arg)
	c:=exec.Command(exe,arg)
	r,err:=c.Output()
	return r,err
}

func (sniffer *Ansible)SniiferServer(servers []*prometheus.Server)(error){
	ip_server_map:=map[string]*prometheus.Server{}
	ip_arr:=[]string{}
	for _,s:=range servers {
		fmt.Println(s.NetPorts)
		if len(s.NetPorts) >0{
			fmt.Println("got 1")
			for _,n:=range s.NetPorts{
				fmt.Println("got 2")
				if n.Ipv4Int !=nil{
					fmt.Println("got 3")
					ipv4_str:=net.Ipv4Uint32ConverString(s.NetPorts[0].Ipv4Int.(uint32))
					ip_server_map[ipv4_str]=s
					ip_arr=append(ip_arr,ipv4_str)
					break
				}
			}
		}
	}
	fmt.Printf("%#v \n",ip_server_map)
	//ip_device_id_map:=map[string]int{"192.168.33.81":1}
	ar,err:=RunAnsibleCommand(sniffer.Exe,strings.Join(ip_arr,`,`))
	if err!=nil{
		return err
	}
	json, err := simplejson.NewJson(ar)
	success_servers,err:=json.Get("contacted").Map()
	failed_servers,err:=json.Get("dark").Map()
	if err!=nil{
		return err
	}
	for ip,_ :=range failed_servers{
		log.Warn("sniff ",ip ," failed:"," connect failed")
	}
	for ip,infomation :=range success_servers{
		//var server prometheus.Server
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
			//server=prometheus.Server{Serial:serial,Hostname:hostname,Os:os,Release:release,Memsize:memsize,NetPorts:netPorts}
		}
		ip_server_map[ip].Serial=serial
		ip_server_map[ip].Hostname=hostname
		ip_server_map[ip].Os=os
		ip_server_map[ip].Release=release
		ip_server_map[ip].Memsize=memsize
		ip_server_map[ip].NetPorts=netPorts
		err:=ip_server_map[ip].UpdateServer()
		if err!=nil{
			log.Warn("sniff ",ip," failed :",err)
		}else{
			log.Debug("sniff ",ip," success")
		}
	}
	return nil
}