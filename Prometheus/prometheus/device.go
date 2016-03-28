package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

//func GetDeviceType() (interface{}, error) {
//	r, err := PROMETHEUS.dbobj.GetDeviceType()
//	return r, err
//}

func GetDeviceModel(args ...[]byte) (interface{}, error) {
	var id int
	var name string
	var _type string
	cur, err := PROMETHEUS.dbobj.Find("device_model", [][]byte{[]byte(`device_model_id`), []byte(`device_model_name`), []byte(`device_type`)}, args, &id, &name, &_type)
	r := []global.DeviceModel{}
	for cur.Fetch() {
		r = append(r, global.DeviceModel{id, name, _type})
	}
	return r, err
}
