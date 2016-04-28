package prometheus

import (
	"database/sql"
	"fmt"
	//"github.com/luckykris/Cronus/Hephaestus/net"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetNetPort(args ...string) (interface{}, error) {
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
func (device *Device) GetNetPort(id ...int) (interface{}, error) {
	conditions := []string{}
	conditions = append(conditions, fmt.Sprintf("device_id=%d", device.DeviceId))
	if len(id) > 0 {
		conditions = append(conditions, fmt.Sprintf("netPort_id=%d", id[0]))
	}
	return GetNetPort(conditions...)
}
