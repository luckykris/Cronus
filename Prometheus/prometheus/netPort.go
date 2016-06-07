package prometheus

import (
	"database/sql"
	"fmt"
	//"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetNetPort(args ...string) ([]NetPort, error) {
	var netPort_id int
	var mac sql.NullString
	var mac_i interface{}
	var ipv4_int sql.NullInt64
	var ipv4_int_i interface{}
	var netPort_type string
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEnetPort, []string{`netPort_id`, `mac`, `ipv4_int`, `netPort_type`}, args, &netPort_id, &mac, &ipv4_int, &netPort_type)
	r := []NetPort{}
	for cur.Fetch() {
		if !mac.Valid {
			mac_i = nil
		} else {
			mac_i = mac.String
		}
		if !ipv4_int.Valid {
			ipv4_int_i = nil
		} else {
			ipv4_int_i = uint32(ipv4_int.Int64)
		}
		r = append(r, NetPort{netPort_id, mac_i, ipv4_int_i, netPort_type})
	}
	return r, err
}

func AddNetPort(device_id int,netPort NetPort)(error) {
	return PROMETHEUS.dbobj.Add(global.TABLEnetPort, []string{`mac`, `ipv4_int`, `device_id`,`netPort_type`}, [][]interface{}{[]interface{}{netPort.Mac,netPort.Ipv4Int,device_id,netPort.Type}})
}

func (device *Device) GetNetPort(id ...int) ([]NetPort, error) {
	conditions := []string{}
	conditions = append(conditions, fmt.Sprintf("device_id=%d", device.DeviceId))
	if len(id) > 0 {
		conditions = append(conditions, fmt.Sprintf("netPort_id=%d", id[0]))
	}
	return GetNetPort(conditions...)
}

func (device *Device) AddNetPort(netPort NetPort) (error) {
	return AddNetPort(device.DeviceId,netPort)
}

func (netPort *NetPort) UpdateNetPort() (error) {
	c := fmt.Sprintf("netPort_id = %d", netPort.NetPortId)
	return PROMETHEUS.dbobj.Update(global.TABLEnetPort,[]string{c}, []string{`mac`, `ipv4_int`,`netPort_type`}, []interface{}{netPort.Mac,netPort.Ipv4Int,netPort.Type})
}

func  (device *Device) DeleteNetPort(netPort *NetPort) (error) {
	c1 := fmt.Sprintf("netPort_id = %d", netPort.NetPortId)
	c2 := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Delete(global.TABLEnetPort, []string{c1,c2})
}