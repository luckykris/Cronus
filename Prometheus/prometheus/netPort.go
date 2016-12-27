package prometheus

import (
	//"database/sql"
	"fmt"
	"time"
	//"strings"
	//"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Hephaestus/go/net"
)

func (device *Device) GetNetPort()[]NetPort{
	defer device.RUnlock()
	device.RLock()
	return device.NetPorts
}

func (device *Device) AddNetPort(netPort NetPort)(err error){
	defer device.Unlock()
	device.Lock()
	err=device.AddNetPortViaDB(netPort)
	if err==nil {
		device.NetPorts=append(device.NetPorts,netPort)
	}
	return err
}
func (device *Device) DeleteNetPort(netPort NetPort)(err error){
	defer device.Unlock()
	device.Lock()
	for i,n:=range device.NetPorts{
		if n.Ipv4==netPort.Ipv4{
			err=device.DeleteNetPortViaDB(netPort)
			if err==nil{
				device.NetPorts=append(device.NetPorts[:i],device.NetPorts[i + 1:]...)
			}
			return err
		}
	}
	return global.ERROR_resource_notexist
}

func (device *Device) AddNetPortViaDB(netPort NetPort)(error) {
	items:=[]string{`device_id`,
					`mac`, 
				    `ipv4_int`, 
				    `netPort_type`,
				    `ctime`,
					}
	ipv4_int,err:=net.Ipv4StringConverUint32(netPort.Ipv4)
	if err!=nil{return err}
	if !net.Ipv4UintAvaliable(ipv4_int,uint(netPort.Mask)){return err}
	values:=[][]interface{}{[]interface{}{device.Get_DeviceId(),
										  netPort.Mac,
										  ipv4_int,
										  netPort.NetPortType,
										  time.Now().Unix(),
										  }}			    
	return PROMETHEUS.dbobj.Add(global.TABLEnetPort, items, values)
}
func  (device *Device) DeleteNetPortViaDB(netPort NetPort) (error) {
	ipv4_int,err:=net.Ipv4StringConverUint32(netPort.Ipv4)
	if err!=nil{return err}
	c1 := fmt.Sprintf("ipv4_int = %d", ipv4_int)
	c2 := fmt.Sprintf("device_id = %d", device.Get_DeviceId())
	return PROMETHEUS.dbobj.Delete(global.TABLEnetPort, []string{c1,c2})
}