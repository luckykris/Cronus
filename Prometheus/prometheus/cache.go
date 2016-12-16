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
		SERVER string="server"
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
		cache_and_index_one_device(s[i],SERVER)
	}
	return err
}

func FlushDeviceCache(device interface{},device_type string){
	switch device_type{
	case SERVER:
		old_device,ok:=DEVICE_INDEX_ID[SERVER][device.(*Server).Device.DeviceId]
		if ok{
			DEVICE_CACHE.Container.Remove(old_device)
		}
		cache_and_index_one_device(device,SERVER)
	}
}
func cache_and_index_one_device(device interface{},device_type string){
	defer DEVICE_CACHE.Unlock()
	DEVICE_CACHE.Lock()
	tmp_e:=DEVICE_CACHE.Container.PushBack(device)
	switch device_type{
	case SERVER:
			DEVICE_INDEX_ID[SERVER][device.(*Server).Device.DeviceId]=tmp_e
			DEVICE_INDEX_NAME[SERVER][device.(*Server).Device.DeviceName]=tmp_e	
	}
}
