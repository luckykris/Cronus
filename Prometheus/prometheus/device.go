package prometheus

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetDevice(id ...int) ([]*Device, error) {
	var device_id int
	var device_name string
	var device_type string
	var father_device_id sql.NullInt64
	var father_device_id_i interface{}
	conditions:=[]string{}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("device_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	r := []*Device{}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdevice,nil, []string{`device_id`, `device_name`, `device_type`, `father_device_id`}, conditions, &device_id, &device_name, &device_type, &father_device_id)
	if err != nil {
		return r, err
	}
	for cur.Fetch() {
		if !father_device_id.Valid {
			father_device_id_i = nil
		} else {
			father_device_id_i = father_device_id.Int64
		}
		r = append(r, &Device{device_id, device_name, device_type, father_device_id_i})
	}
	return r, err
}

func AddDevice(device *Device) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdevice, []string{`device_name`,`device_type`, `father_device_id`}, [][]interface{}{[]interface{}{device.DeviceName,device.DeviceType,device.FatherDeviceId}})
}

func (device *Device)DeleteDevice() error {
	c := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Delete(global.TABLEdevice, []string{c})
}

func (device *Device)UpdateDevice() error {
	c := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Update(global.TABLEdevice, []string{c}, []string{`device_name`,`device_type`, `father_device_id`}, []interface{}{device.DeviceName,device.DeviceType,device.FatherDeviceId})
}
