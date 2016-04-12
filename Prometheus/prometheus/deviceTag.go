package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetDeviceTag(args ...string) (interface{}, error) {
	var device_id int
	var tag_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceTag, []string{`device_id`, `tag_id`}, args, &device_id, &tag_id)
	r := []global.DeviceTag{}
	for cur.Fetch() {
		r = append(r, global.DeviceTag{device_id,tag_id})
	}
	return r, err
}
