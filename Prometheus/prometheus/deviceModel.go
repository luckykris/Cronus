package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

//func GetDeviceType() (interface{}, error) {
//	r, err := PROMETHEUS.dbobj.GetDeviceType()
//	return r, err
//}

func GetDeviceModel(args ...string) (interface{}, error) {
	var id int
	var name string
	var _type string
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceModel, []string{`device_model_id`, `device_model_name`, `device_type`}, args, &id, &name, &_type)
	r := []global.DeviceModel{}
	for cur.Fetch() {
		r = append(r, global.DeviceModel{id, name, _type})
	}
	return r, err
}

func AddDeviceModel(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdeviceModel, []string{`device_model_name`, `device_type`}, values)
}

func DeleteDeviceModel(id int) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceModel, []string{c})
}

func UpdateDeviceModel(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLEdeviceModel, []string{c}, cloumns, values)
}
