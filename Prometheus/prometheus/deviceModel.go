package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)



func GetDeviceModel(id ...int) ([]*DeviceModel,error) {
	deviceModels:=[]*DeviceModel{}
	if len(id) !=0 {
		for _,v:=range id{
			deviceModels=append(deviceModels,PROMETHEUS.DeviceModelMapId[v])
		}
		return deviceModels,nil
	}else{
		for _,v:=range PROMETHEUS.DeviceModelMapId{
			deviceModels=append(deviceModels,v)
		}
		return deviceModels,nil
	}
}


func CacheDeviceModel(name interface{},id ...int) (error) {
	var device_model_id int
	var device_name string
	var _type string
	conditions:=[]string{}
	if name !=nil{
		conditions=append(conditions,fmt.Sprintf(`device_model_name = '%s'`,name.(string)))
	}
	if len(id) >0 {
		id_str_ls:=[]string{}
		for _,v :=range id{
			id_str_ls=append(id_str_ls,fmt.Sprintf("%d",v))
		}
		conditions=append( conditions,fmt.Sprintf(`device_model_id IN (%s)`,strings.Join(id_str_ls,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceModel, nil,[]string{`device_model_id`, `device_model_name`, `device_type`}, conditions, &device_model_id, &device_name, &_type)
	for cur.Fetch() {
		deviceModel := new(DeviceModel)
		deviceModel.DeviceModelId=device_model_id
		deviceModel.DeviceModelName=device_name
		deviceModel.DeviceType=_type
		PROMETHEUS.DeviceModelMapId[device_model_id]=deviceModel
	}
	return err
}

func AddDeviceModel(device_model_name ,device_type string) error {
	err:=PROMETHEUS.dbobj.Add(global.TABLEdeviceModel, []string{`device_model_name`, `device_type`}, [][]interface{}{[]interface{}{device_model_name,device_type}})
	if err!=nil{
		return err
	}
	return CacheDeviceModel(device_model_name)
}

func DeleteDeviceModel(id int) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceModel, []string{c})
}

func UpdateDeviceModel(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("device_model_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLEdeviceModel, []string{c}, cloumns, values)
}
