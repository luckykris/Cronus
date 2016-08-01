package prometheus

import (
	//"database/sql"
	"time"
	"crypto/md5"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	//log "github.com/Sirupsen/logrus"
)


func GetServer(name interface{},id ...int) ([]*Server, error){
	servers:=[]*Server{}
	if len(id) !=0 {
		for _,v:=range id{
			servers=append(servers,PROMETHEUS.ServerMapId[v])
		}
		return servers,nil
	}else{
		for _,v:=range PROMETHEUS.ServerMapId{
			servers=append(servers,v)
		}
		return servers,nil
	}
}
func CacheServer(name interface{},id ...int) (error) {
	var device_id int
	var deviceName string
	var deviceModel *DeviceModel
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
	if err!=nil{
		return err
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
		return  err
	}
	for _,device:=range devices {
		device_map[device.DeviceId] = device
	}
	for cur.Fetch() {
		if _,ok:=device_map[device_id];!ok{
			return fmt.Errorf("device data fatal!")
		}
		deviceName=device_map[device_id].DeviceName
		deviceModel=device_map[device_id].DeviceModel
		fatherDeviceId=device_map[device_id].FatherDeviceId
		server:=new(Server)
		server.Device.NetPorts=device_map[device_id].NetPorts
		server.Device.DeviceId=device_id
		server.Device.DeviceName=deviceName
		server.Device.DeviceModel=deviceModel
		server.Device.FatherDeviceId=fatherDeviceId
		server.Serial=serial
		server.Hostname=hostname
		server.Memsize=memsize
		server.Os=os
		server.Release=release
		server.LastChangeTime=last_change_time
		server.Checksum=checksum
		PROMETHEUS.ServerMapId[server.Device.DeviceId]=server
	}
	return err
}



func AddServer(device_name string,device_model_id int) error {
	err:=AddDevice(device_name,device_model_id,nil)
	if err!=nil{
		return err
	}
	devices,err:=GetDevice(device_name)
	if err!=nil{
		return err
	}
	if len(devices)!=1{
		return fmt.Errorf("amazing critical error!")
	}
	err=PROMETHEUS.dbobj.Add(global.TABLEserver, []string{`device_id`,`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, [][]interface{}{[]interface{}{devices[0].DeviceId,"Unknow","Unknow",0,"Unknow",0,0,"Never"}})
	if err!=nil{
		return err
	}else{
		server:=new(Server)
		server.Device.DeviceId=devices[0].DeviceId
		server.Device.DeviceName=devices[0].DeviceName
		server.Device.DeviceModel=devices[0].DeviceModel
		server.Device.FatherDeviceId=devices[0].FatherDeviceId
		server.Serial="Unknow"
		server.Hostname="Unknow"
		server.Memsize=0
		server.Os="Unknow"
		server.Release=0
		server.LastChangeTime=0
		server.Checksum="Never"
		PROMETHEUS.ServerMapId[devices[0].DeviceId]=server
		return nil
	}
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
	err:=PROMETHEUS.dbobj.Update(global.TABLEserver, conditions, []string{`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, []interface{}{server.Serial,server.Hostname,server.Memsize,server.Os,server.Release,server.LastChangeTime,checksum})
	if err!=nil{
		return err
	}
	err:=CacheServer(nil,server.Device.DeviceId)
	if err!=nil{
		return err
	}
	return nil
}


func(server *Server)ComputSum()string{
	s:=fmt.Sprintf("%s%s%d%s%f",server.Serial,server.Hostname,server.Memsize,server.Os,server.Release)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range server.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}

