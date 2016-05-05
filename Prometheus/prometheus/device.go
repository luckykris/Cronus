package prometheus

import (
	"database/sql"
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetDevice(args ...string) ([]Device, error) {
	var device_id int
	var device_name string
	var device_model_id int
	var father_device_id sql.NullInt64
	var father_device_id_i interface{}
	r := []Device{}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdevice, []string{`device_id`, `device_name`, `device_model_id`, `father_device_id`}, args, &device_id, &device_name, &device_model_id, &father_device_id)
	if err != nil {
		return r, err
	}
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

func AddDevice(device *Device) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdevice, []string{`device_name`,`device_model_id`, `father_device_id`}, [][]interface{}{[]interface{}{device.DeviceName,device.DeviceModelId,device.FatherDeviceId}})
}

func DeleteDevice(id int) error {
	c := fmt.Sprintf("device_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEdevice, []string{c})
}

func (device *Device)UpdateDevice() error {
	c := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Update(global.TABLEdevice, []string{c}, []string{`device_name`,`device_model_id`, `father_device_id`}, []interface{}{device.DeviceName,device.DeviceModelId,device.FatherDeviceId})
}
