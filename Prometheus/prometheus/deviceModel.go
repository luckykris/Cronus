package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)




func GetOneDeviceModel(device_model_id interface{},device_model_name interface{})(*DeviceModel,error){
	if device_model_id !=nil{
		s,ok:=DEVICEMODEL_INDEX_ID[device_model_id.(int)]
		if ok {
			return s.Value.(*DeviceModel),nil
		}
		return &DeviceModel{},global.ERROR_resource_notexist
	}else if device_model_name !=nil{
		s,ok:=DEVICEMODEL_INDEX_NAME[device_model_name.(string)]
		if ok {
			return s.Value.(*DeviceModel),nil
		}
		return &DeviceModel{},global.ERROR_resource_notexist
	}
	return &DeviceModel{},global.ERROR_resource_notexist
}
func GetDeviceModel()([]*DeviceModel,error){
	r:=[]*DeviceModel{}
	for _,v:=range DEVICEMODEL_INDEX_ID{
		r=append(r,v.Value.(*DeviceModel))
	}
	return r,nil
}
func AddDeviceModel(deviceModel *DeviceModel)error{
	err:=AddDeviceModelViaDB(deviceModel)
	if err==nil {
		create_cache_and_index(deviceModel)
	}
	return err
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
	items:=[]string{`device_model_id`, 
		            `device_model_name`, 
		            `device_type`,
		            `u`,
		            `half_full`}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceModel, nil,items, conditions, &device_model_id,
																					 &device_model_name, 
																					 &device_type ,
																					 &u,
																					 &half_full)
	if err!=nil{
		return
	}
	for cur.Fetch() {
		deviceModel := new(DeviceModel)
		deviceModel.DeviceModelId=device_model_id
		deviceModel.DeviceModelName=device_model_name
		deviceModel.DeviceType=device_type
		deviceModel.U=u
		deviceModel.HALF_FULL=half_full
		result=append(result,deviceModel)
	}
	return 
}

func AddDeviceModelViaDB(deviceModel *DeviceModel)(err error) {
	_,err=GetOneDeviceModel(nil,deviceModel.DeviceModelName)
	if err==nil{
		return global.ERROR_resource_duplicate
	}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{
		return err
	}
	defer func(){
		if err!=nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	rows:=[][]interface{}{[]interface{}{deviceModel.DeviceModelName,
										deviceModel.DeviceType,
										deviceModel.U,
										deviceModel.HALF_FULL,
										}}
	err=tx.Add(global.TABLEdeviceModel,[]string{`device_model_name`,
									  	  `device_type`,
									  	  `u`,
									  	  `half_full`,
									      },rows)
	if err!=nil{
		return err
	}
	var device_model_id int
	conditions:=[]string{fmt.Sprintf("device_model_name='%s'",deviceModel.DeviceModelName)}
	items:=[]string{"device_model_id"}
	cur,err:=tx.Get(global.TABLEdeviceModel, nil,items, conditions,  
					&device_model_id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	deviceModel.DeviceModelId=device_model_id
	cur.Close()
	return err
}

func (self *DeviceModel)Delete()(err error){
	defer self.Unlock()
	self.Lock()
	err=self.DeleteViaDB()
	if err!=nil{
		return 
	}
	drop_cache_and_index(self)
	return 
}
func (self *DeviceModel)DeleteViaDB()error{
	conditions:=[]string{}
	conditions=append(conditions,fmt.Sprintf("device_model_id=%d",self.DeviceModelId))
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceModel,conditions)
}

func (self *DeviceModel)Get_DeviceType()string{
	return self.DeviceType
}