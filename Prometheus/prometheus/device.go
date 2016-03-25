package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

//func GetDeviceType() (interface{}, error) {
//	r, err := PROMETHEUS.dbobj.GetDeviceType()
//	return r, err
//}

func GetDeviceType() (interface{}, error) {
	var id int
	var name string
	cur, err := PROMETHEUS.dbobj.Find("device_model", [][]byte{[]byte(`device_model_id`), []byte(`device_model_name`)}, &id, &name)
	r := []global.DeviceType{}
	for cur.Fetch() {
		r = append(r, global.DeviceType{id, name})
	}
	return r, err
}
