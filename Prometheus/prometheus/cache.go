package prometheus
import (
	//"github.com/luckykris/Cronus/Prometheus/global"
	//"database/sql"
	"container/list"
)


func Cache_Index_Init()(error) {
	var err error=nil
	//init index
	DEVICEMODEL_INDEX_ID=map[int]*list.Element{}
	DEVICE_INDEX_ID=map[string]map[int]*list.Element{}
	DEVICE_INDEX_ID["server"]=map[int]*list.Element{}
	
	//init cache list
	DEVICE_CACHE=list.New()
	DEVICEMODEL_CACHE=list.New()
	m,err:=GetDeviceModelFromDB([]string{},[]string{},[]int{})
	if err!=nil{
		return err
	}
	for i:=range m{
		DEVICEMODEL_INDEX_ID[m[i].DeviceModelId]=DEVICEMODEL_CACHE.PushBack(m[i])
	}
	s,err:=GetServerFromDB([]int{},[]string{},[]int{} ,[]uint8{})
	if err!=nil{
		return err
	}
	for i:=range s{
		DEVICE_INDEX_ID["server"][s[i].Device.DeviceId]=DEVICE_CACHE.PushBack(s[i])
	}
	return err
}

