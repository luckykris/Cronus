package prometheus

import (
	//"database/sql"
	"time"
	"crypto/md5"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)



func GetServer(name interface{},id ...int) ([]*Server, error) {
	var device_id int
	var deviceName string
	var deviceType string
	var fatherDeviceId interface{}
	var serial string
	var hostname string
	var memsize int
	var os string
	var release float64
	var last_change_time int64
	var checksum string
	conditions:=[]string{}
	devices,err:=GetDevice(name,id...)
	device_map:=map[int]*Device{}
	servers:=[]*Server{}
	if err!=nil{
		return []*Server{},err
	}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("device_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEserver, nil, []string{`device_id`,`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, conditions, &device_id, &serial, &hostname, &memsize,&os,&release,&last_change_time,&checksum)
	if err != nil {
		return servers, err
	}
	for _,device:=range devices {
		device_map[device.DeviceId] = device
	}
	for cur.Fetch() {
		deviceName=device_map[device_id].DeviceName
		deviceType=device_map[device_id].DeviceType
		fatherDeviceId=device_map[device_id].FatherDeviceId
		server:=new(Server)
		server.Device.DeviceId=device_id
		server.Device.DeviceName=deviceName
		server.Device.DeviceType=deviceType
		server.Device.FatherDeviceId=fatherDeviceId
		server.Serial=serial
		server.Hostname=hostname
		server.Memsize=memsize
		server.Os=os
		server.Release=release
		server.LastChangeTime=last_change_time
		server.Checksum=checksum

		//server:=&Server{
		//			  Serial:serial,
		//			  Hostname :hostname,
		//			  Memsize: memsize,
		//			  Os:os,
		//			  Release :release,
		//			  LastChangeTime: last_change_time,
		//			  Checksum :checksum}
		//server.Init(device_id,deviceName,deviceType,fatherDeviceId)
		netPorts,err:=server.GetNetPort()
		if err!=nil{
			log.Error("prometheus get netPort failed:",err.Error())
		}
		server.NetPorts=netPorts
		servers=append(servers,server)
	}
	return servers, err
}

func NewServer(deviceName string)*Server{
	server:=new(Server)
	server.Device.DeviceName=deviceName
	server.Device.DeviceType="Server"
	return server
}

func (server *Server)AddServer() error {
	err:=server.AddDevice()
	if err!=nil{
		return err
	}
	devices,err:=GetDevice(server.Device.DeviceName)
	if err!=nil{
		return err
	}
	if len(devices)!=1{
		return fmt.Errorf("amazing critical error!")
	}
	return PROMETHEUS.dbobj.Add(global.TABLEserver, []string{`device_id`,`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, [][]interface{}{[]interface{}{devices[0].DeviceId,"Unknow","Unknow",0,"Unknow",0,0,"Never"}})
}
//
func (server *Server)DeleteServer() error {
	return server.DeleteDevice()
}

func (server *Server)UpdateServer() error {
	conditions:=[]string{fmt.Sprintf("device_id = %d" , server.DeviceId)}
	server_pre_ls,err:=GetServer(nil,server.DeviceId)
	if err!=nil{
		return err
	}
	checksum:=server.ComputSum()
	if server_pre_ls[0].Checksum == checksum{
		return nil
	}
	server.LastChangeTime=time.Now().Unix()
	return PROMETHEUS.dbobj.Update(global.TABLEserver, conditions, []string{`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, []interface{}{server.Serial,server.Hostname,server.Memsize,server.Os,server.Release,server.LastChangeTime,checksum})
}


func(server *Server)ComputSum()string{
	s:=fmt.Sprintf("%s%s%d%s%f",server.Serial,server.Hostname,server.Memsize,server.Os,server.Release)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range server.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}