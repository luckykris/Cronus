package prometheus

import (
	//"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)



func GetServer(id ...int) ([]Server, error) {
	var device_id int
	var deviceName string
	var deviceType string
	var fatherDeviceId interface{}
	var serial string
	var hostname string
	var memsize int
	var os string
	var release float64
	var last_change_time int
	var checksum string
	conditions:=[]string{}
	devices,err:=GetDevice(id...)
	device_map:=map[int]Device{}
	servers:=[]Server{}
	if err!=nil{
		return []Server{},err
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
		//server_map[device_id].Serial=serial
		//server_map[device_id].Serial=serial
		//server_map[device_id].Hostname=hostname
		//server_map[device_id].Memsize=memsize
		//server_map[device_id].Os=os
		//server_map[device_id].Release=release
		//server_map[device_id].LastChangeTime=last_change_time
		//server_map[device_id].Checksum=checksum
		server:=Server{
					  Serial:serial,
					  Hostname :hostname,
					  Memsize: memsize,
					  Os:os,
					  Release :release,
					  LastChangeTime: last_change_time,
					  Checksum :checksum}
		server.Init(device_id,deviceName,deviceType,fatherDeviceId)
		servers=append(servers,server)
	}
	return servers, err
}

//func AddDevice(device *Device) error {
//	return PROMETHEUS.dbobj.Add(global.TABLEdevice, []string{`device_name`,`device_type`, `father_device_id`}, [][]interface{}{[]interface{}{device.DeviceName,device.DeviceType,device.FatherDeviceId}})
//}
//
//func (device *Device)DeleteDevice() error {
//	c := fmt.Sprintf("device_id = %d", device.DeviceId)
//	return PROMETHEUS.dbobj.Delete(global.TABLEdevice, []string{c})
//}
//
func (server *Server)UpdateServer() error {

	return PROMETHEUS.dbobj.Update(global.TABLEdevice, []string{c}, []string{`device_name`,`device_type`, `father_device_id`}, []interface{}{device.DeviceName,device.DeviceType,device.FatherDeviceId})
}


func(server *Server)Checksum()string{
	for _,v :=range server.NetPorts{

	}
}