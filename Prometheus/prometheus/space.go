package prometheus

import (
	"fmt"
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


func (device *Device)GetSpace() ([]Space,error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	conditions:=[]string{fmt.Sprintf("device_id=%d",device.DeviceId)}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, conditions, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r,err
}


//func (device *Device)AddSpace(cabinet_id ,start_u int)(error){
//	spaces:=[][]interface{}{}
//	err:=PROMETHEUS.dbobj.Add(global.TABLEspace, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, [][]interface{}{[]interface{}{device.DeviceName,device.DeviceType,device.FatherDeviceId}})
//	return err
//}