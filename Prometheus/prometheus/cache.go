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
var LOCATION_CACHE    =new(Cache)
var IDC_CACHE         =new(Cache)
var CABINET_CACHE     =new(Cache)

var DEVICEMODEL_INDEX_ID 	map[int]*list.Element
var LOCATION_INDEX_ID       map[int]*list.Element
var DEVICE_INDEX_ID 		map[string]map[int]*list.Element
var DEVICE_INDEX_NAME 		map[string]map[string]*list.Element
var IDC_INDEX_ID 		    map[int]*list.Element
var CABINET_INDEX_ID 		map[int]*list.Element

func Cache_Index_Init()(error) {
	start_t:=time.Now().UnixNano()
	defer func(){
			log.Info(fmt.Sprintf("cache init cost:%d ms" ,(time.Now().UnixNano()-start_t)/1000/1000))
		}()
	var err error=nil
	var tmp_name string
	err=init_cache_and_index_Location()    ;tmp_name="Location";   if err!=nil{log.Fatal(fmt.Sprintf("%s cache init failed:%s",tmp_name,err.Error()));return err }else{log.Debug(fmt.Sprintf("%s cache init success",tmp_name))}
	err=init_cache_and_index_Idc()         ;tmp_name="Idc";        if err!=nil{log.Fatal(fmt.Sprintf("%s cache init failed:%s",tmp_name,err.Error()));return err }else{log.Debug(fmt.Sprintf("%s cache init success",tmp_name))}
	err=init_cache_and_index_Cabinet()     ;tmp_name="Cabinet";    if err!=nil{log.Fatal(fmt.Sprintf("%s cache init failed:%s",tmp_name,err.Error()));return err }else{log.Debug(fmt.Sprintf("%s cache init success",tmp_name))}  
	err=init_cache_and_index_DeviceModel() ;tmp_name="DeviceModel";if err!=nil{log.Fatal(fmt.Sprintf("%s cache init failed:%s",tmp_name,err.Error()));return err }else{log.Debug(fmt.Sprintf("%s cache init success",tmp_name))}  
	err=init_cache_and_index_Device()      ;tmp_name="Device";     if err!=nil{log.Fatal(fmt.Sprintf("%s cache init failed:%s",tmp_name,err.Error()));return err }else{log.Debug(fmt.Sprintf("%s cache init success",tmp_name))}    
	return err
}
func init_cache_and_index_Cabinet()error{
	CABINET_INDEX_ID=map[int]*list.Element{}
	CABINET_CACHE.Container=list.New()
	l,err:=GetCabinetViaDB([]int{},[]string{})
	if err!=nil{
		return err
	}
	for i:=range l{
		CABINET_INDEX_ID[l[i].CabinetId]=CABINET_CACHE.Container.PushBack(l[i])
	}
	return nil
}
func init_cache_and_index_Idc()error{
	IDC_INDEX_ID=map[int]*list.Element{}
	IDC_CACHE.Container=list.New()
	l,err:=GetIdcViaDB([]int{},[]string{})
	if err!=nil{
		return err
	}
	for i:=range l{
		IDC_INDEX_ID[l[i].IdcId]=IDC_CACHE.Container.PushBack(l[i])
	}
	return nil
}
func init_cache_and_index_Location()error{
	LOCATION_INDEX_ID=map[int]*list.Element{}
	LOCATION_CACHE.Container=list.New()
	l,err:=GetLocationViaDB([]int{},[]string{})
	if err!=nil{
		return err
	}
	for i:=range l{
		LOCATION_INDEX_ID[l[i].LocationId]=LOCATION_CACHE.Container.PushBack(l[i])
	}
	return nil
}
func init_cache_and_index_DeviceModel()error{
	DEVICEMODEL_INDEX_ID=map[int]*list.Element{}
	DEVICEMODEL_CACHE.Container=list.New()
	m,err:=GetDeviceModelViaDB([]string{},[]string{},[]int{})
	if err!=nil{
		return err
	}
	for i:=range m{
		DEVICEMODEL_INDEX_ID[m[i].DeviceModelId]=DEVICEMODEL_CACHE.Container.PushBack(m[i])
	}
	return nil
}
func init_cache_and_index_Device()error{
	DEVICE_INDEX_ID=map[string]map[int]*list.Element{}
	DEVICE_INDEX_NAME=map[string]map[string]*list.Element{}
	DEVICE_INDEX_ID[SERVER]=map[int]*list.Element{}
	DEVICE_INDEX_NAME[SERVER]=map[string]*list.Element{}
	DEVICE_CACHE.Container=list.New()
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