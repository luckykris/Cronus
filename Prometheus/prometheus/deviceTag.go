package prometheus

import (
	"fmt"
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
func AddDeviceTag(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdeviceTag, []string{`device_id`, `tag_id`}, values)
}

func DeleteDeviceTag(device_id , tag_id int) error {
	device_id_c := fmt.Sprintf("device_id = %d", device_id)
	tag_id_c:= fmt.Sprintf("tag_id = %d", tag_id)
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceTag, []string{device_id_c,tag_id_c})
}
