package prometheus

import (
	//"database/sql"
	"fmt"
	//"strings"
	//"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
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
	if err==nil && ifcache(){
		device.NetPorts=append(device.NetPorts,netPort)
	}
	return err
}

func (device *Device) AddNetPortViaDB(netPort NetPort)(error) {
	items:=[]string{`device_id`,
					`mac`, 
				    `ipv4_int`, 
				    `netPort_type`}
	values:=[][]interface{}{[]interface{}{device.DeviceId,
										  netPort.Mac,
										  netPort.Ipv4,
										  netPort.NetPortType}}			    
	return PROMETHEUS.dbobj.Add(global.TABLEnetPort, items, values)
}





func  (device *Device) DeleteNetPort(netPort *NetPort) (error) {
	c1 := fmt.Sprintf("ipv4_int = %d", netPort.Ipv4)
	c2 := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Delete(global.TABLEnetPort, []string{c1,c2})
}