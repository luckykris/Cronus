package prometheus

import (
	"fmt"
	"database/sql"
	"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetNetPort(args ...string) (interface{}, error) {
	var mac sql.NullString
	var mac_i interface{}
	var ipv4_int sql.NullInt64
	var ipv4_int_i interface{}
	var _type string
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEnetPort, []string{`mac`, `ipv4_int`, `type`}, args, &mac, &ipv4_int, &_type)
	r := []global.NetPort{}
	for cur.Fetch() {
		if !mac.Valid {
			mac_i = nil
		} else {
			mac_i = mac.String
		}
		if !ipv4_int.Valid {
			ipv4_int_i = nil
		} else {
			ipv4_int_i = net.Ipv4Uint32ConverString(uint32(ipv4_int.Int64))
		}
		r = append(r, global.NetPort{mac_i, ipv4_int_i, device_id, _type})
	}
	return r, err
}
func (device *global.Device)GetNetPort(id ...int)(interface{},error){
	conditions:=[]string{}
	conditions=append(conditions,fmt.Sprintf("device_id=%d",device.DeviceId))
	if len(id)>0{
		conditions=append(conditions,fmt.Sprintf("ipv4_int" ,id[0]))
	}
	return GetNetPort(conditions ...)
}
