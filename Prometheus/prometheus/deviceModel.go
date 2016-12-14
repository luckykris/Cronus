package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)





func GetDeviceModelFromDB(names []string,device_types []string,device_model_ids []int) (result []*DeviceModel,err error) {
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



