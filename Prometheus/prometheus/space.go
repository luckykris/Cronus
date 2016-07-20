package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetSpace(args ...string) (interface{}, error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, args, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r, err
}
