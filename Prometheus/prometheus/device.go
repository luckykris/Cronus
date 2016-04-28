package prometheus

import (
	"database/sql"
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetDevice(args ...string) (interface{}, error) {
	var device_id int
	var device_name string
	var device_model_id int
	var father_device_id sql.NullInt64
	var father_device_id_i interface{}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdevice, []string{`device_id`, `device_name`, `device_model_id`, `father_device_id`}, args, &device_id, &device_name, &device_model_id, &father_device_id)
	if err != nil {
		return nil, err
	}
	r := []Device{}
	for cur.Fetch() {
		if !father_device_id.Valid {
			father_device_id_i = nil
		} else {
			father_device_id_i = father_device_id.Int64
		}
		r = append(r, Device{device_id, device_name, device_model_id, father_device_id_i})
	}
	return r, err
}

func AddDevice(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdeviceModel, []string{`device_model_name`, `device_type`}, values)
}

func DeleteDevice(id int) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceModel, []string{c})
}

func UpdateDevice(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLEdeviceModel, []string{c}, cloumns, values)
}
