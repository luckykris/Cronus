package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)




func GetDeviceModel(device_model_id interface{})(*DeviceModel,error){
	if device_model_id !=nil{
		s,ok:=DEVICEMODEL_INDEX_ID[device_model_id.(int)]
		if ok {
			return s.Value.(*DeviceModel),nil
		}
		return &DeviceModel{},global.ERROR_resource_notexist
	}
	return &DeviceModel{},global.ERROR_resource_notexist
}
func GetDeviceModel_List()([]*DeviceModel){
	r:=[]*DeviceModel{}
	for _,v:=range DEVICEMODEL_INDEX_ID{
		r=append(r,v.Value.(*DeviceModel))
	}
	return r
}

func GetDeviceModelViaDB(names []string,device_types []string,device_model_ids []int) (result []*DeviceModel,err error) {
	var half_full string
	var u uint8
	var device_model_id int
	var device_model_name string
	var device_type string
	conditions:=[]string{}

	err = nil
	result=[]*DeviceModel{}
	if len(device_types) >0{
		conditions=append(conditions,fmt.Sprintf(`device_type IN ('%s')`,strings.Join(device_types,"','")))
	}
	if len(names) >0 {
		conditions=append( conditions,fmt.Sprintf(`device_model_name IN ('%s')`,strings.Join(names,"','")))
	}
	if len(device_model_ids) >0 {
		conditions=append( conditions,fmt.Sprintf(`device_model_id IN (%s)`,int_join(device_model_ids,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceModel, nil,[]string{`device_model_id`, `device_model_name`, `device_type`,`u`,`half_full`}, conditions, &device_model_id, &device_model_name, &device_type ,&u,&half_full)
	if err!=nil{
		return
	}
	for cur.Fetch() {
		deviceModel := new(DeviceModel)
		deviceModel.DeviceModelId=device_model_id
		deviceModel.DeviceModelName=device_model_name
		deviceModel.DeviceType=device_type
		deviceModel.U=u
		result=append(result,deviceModel)
	}
	return 
}



func (self *DeviceModel)Get_DeviceType()string{
	return self.DeviceType
}