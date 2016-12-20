package prometheus
import (
	//"github.com/luckykris/Cronus/Prometheus/global"
	//"database/sql"
	log "github.com/Sirupsen/logrus"
	"time"
	"container/list"
	"fmt"
	"sync"
)
const(
		SERVER string="Server"
)
type Cache struct{
	Container *list.List
	sync.RWMutex
}
var DEVICE_CACHE 	  =new(Cache)
var DEVICEMODEL_CACHE =new(Cache)
var DEVICEMODEL_INDEX_ID 	map[int]*list.Element
var DEVICE_INDEX_ID 		map[string]map[int]*list.Element
var DEVICE_INDEX_NAME 		map[string]map[string]*list.Element

func Cache_Index_Init()(error) {
	start_t:=time.Now().UnixNano()
	defer func(){
			log.Debug(fmt.Sprintf("cache init cost:%d ms" ,(time.Now().UnixNano()-start_t)/1000/1000))
		}()
	var err error=nil
	//init index
	DEVICEMODEL_INDEX_ID=map[int]*list.Element{}
	DEVICE_INDEX_ID=map[string]map[int]*list.Element{}
	DEVICE_INDEX_NAME=map[string]map[string]*list.Element{}
	DEVICE_INDEX_ID[SERVER]=map[int]*list.Element{}
	DEVICE_INDEX_NAME[SERVER]=map[string]*list.Element{}
	//init cache list
	DEVICE_CACHE.Container=list.New()
	DEVICEMODEL_CACHE.Container=list.New()
	m,err:=GetDeviceModelViaDB([]string{},[]string{},[]int{})
	if err!=nil{
		return err
	}
	for i:=range m{
		DEVICEMODEL_INDEX_ID[m[i].DeviceModelId]=DEVICEMODEL_CACHE.Container.PushBack(m[i])
	}
	s,err:=GetServerViaDB([]int{},[]string{},[]int{} ,[]uint8{})
	if err!=nil{
		return err
	}
	for i:=range s{
		create_device_index(s[i])
	}
	return err
}

func FlushDeviceCache(device Device_i){
	drop_device_cache(device)
	create_device_index(device)
}
func DropDeviceCache(device Device_i){
	drop_device_cache(device )
	drop_device_index(device )
}
func drop_device_cache(device Device_i){
	switch device.Get_DeviceModel().Get_DeviceType(){
	case SERVER:
		old_device,ok:=DEVICE_INDEX_ID[SERVER][device.Get_DeviceId()]
		if ok{
			DEVICE_CACHE.Container.Remove(old_device)
		}
	default:
		panic("device type not support")	
	}
}
func create_device_index(device Device_i){
	defer DEVICE_CACHE.Unlock()
	DEVICE_CACHE.Lock()
	switch device.Get_DeviceModel().Get_DeviceType(){
	case SERVER:
		tmp_e:=DEVICE_CACHE.Container.PushBack(device)
		DEVICE_INDEX_ID[SERVER][device.Get_DeviceId()]=tmp_e
		DEVICE_INDEX_NAME[SERVER][device.Get_DeviceName()]=tmp_e
	default:
		panic("device type not support")	
	}
}
func drop_device_index(device Device_i){
	switch device.Get_DeviceModel().Get_DeviceType(){
	case SERVER:
			delete(DEVICE_INDEX_ID[SERVER],device.(*Server).Get_DeviceId())
			delete(DEVICE_INDEX_NAME[SERVER],device.(*Server).Get_DeviceName())
	default:
		panic("device type not support")	
	}
}