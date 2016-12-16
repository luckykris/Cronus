package prometheus

import (
	//"database/sql"
	"fmt"
	//"strings"
	//"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
)


func AddNetPort(device_id int,netPort NetPort)(error) {
	return PROMETHEUS.dbobj.Add(global.TABLEnetPort, []string{`mac`, `ipv4_int`, `device_id`,`netPort_type`}, [][]interface{}{[]interface{}{netPort.Mac,netPort.Ipv4Int,netPort.NetPortType}})
}



func (device *Device) AddNetPort(netPort NetPort) (error) {
	//err:=AddNetPort(device.DeviceId,netPort)
	//if err!=nil{
	//	return err
	//}else{
	//	LoadServer(device.DeviceId)
	//	return nil
	//}
	return AddNetPort(device.DeviceId,netPort)
}


func  (device *Device) DeleteNetPort(netPort *NetPort) (error) {
	c1 := fmt.Sprintf("ipv4_int = %d", netPort.Ipv4Int)
	c2 := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Delete(global.TABLEnetPort, []string{c1,c2})
}