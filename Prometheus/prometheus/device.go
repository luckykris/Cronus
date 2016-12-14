		package prometheus

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)

//func GetDeviceModel(id ...int)([]*DeviceModel,error) {
//	deviceModels:=[]*DeviceModel{}
//	if len(id) !=0 {
//		for _,v:=range id{
//			deviceModels=append(deviceModels,PROMETHEUS.DeviceModelMapId[v])
//		}
//		return deviceModels,nil
//	}else{
//		for _,v:=range PROMETHEUS.DeviceModelMapId{
//			deviceModels=append(deviceModels,v)
//		}
//		return deviceModels,nil
//	}
//}

func GetDeviceFromDB(name interface{},id ...int) ([]*Device, error) {
	var device_id int
	var device_name string
	var father_device_id sql.NullInt64
	var father_device_id_i interface{}
	var device_model_id int
	var ctime uint64
	var group_id int
	var env int
	conditions:=[]string{}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("device_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	if name != nil{
		conditions=append(conditions,fmt.Sprintf("device_name = '%s'",name.(string)))
	}
	r := []*Device{}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdevice,nil, []string{`device_id`, `device_name`,`device_model_id`, `father_device_id`,`ctime`,`group_id`,`env`}, conditions, &device_id, &device_name,&device_model_id, &father_device_id,&ctime,&group_id,&env)
	if err != nil {
		return r, err
	}
	for cur.Fetch() {
		if !father_device_id.Valid {
			father_device_id_i = nil
		} else {
			father_device_id_i = father_device_id.Int64
		}
		device:=new(Device)
		device.DeviceId=device_id
		device.DeviceName=device_name
		device.FatherDeviceId=father_device_id_i
		device.DeviceModel=PROMETHEUS.DeviceModelMapId[device_model_id]
		netPorts,err:=device.GetNetPort()
		if err!=nil{
			log.Error("prometheus get netPort failed:",err.Error())
		}
		device.NetPorts=netPorts
		r = append(r, device)
	}
	return r, err
}


