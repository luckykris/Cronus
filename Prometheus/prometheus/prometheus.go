package prometheus

import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	"os"
	"sync"
)

const (
	HALF string="half"
	FULL string="full"
)
type Device_i interface{
	//GetTag() ([]Tag, error)
	//AddTag(...Tag) error
	//DeleteTag(...Tag) error
	Delete()error
	Get_DeviceId()int
	Get_DeviceName()string
	Get_DeviceModel()*DeviceModel
}

type Device struct {
	sync.RWMutex
	DeviceId       int
	DeviceName     string
	FatherDeviceId interface{}
	Ctime   	   uint64
	GroupId   	   int
	Env	           uint8
	DeviceModel  *DeviceModel
	NetPorts	[]NetPort
	CabinetId	    interface{}
}

type Server struct {
	Serial string
	Hostname string
	Memsize uint32
	Processor uint8
	Os  string
	Release float64
	LastChangeTime uint64
	Checksum string
	Device 
}

type Vm struct {
	Hostname string
	Memsize uint32
	Processor uint8
	Os  string
	Release float64
	LastChangeTime uint64
	Checksum string
	Device 
}

type NetPort struct {
	Mac       interface{}
	Ipv4      string
	NetPortType      string
	Mask	  uint8
}

type Space struct {
	CabinetId int
	DeviceId  interface{}
	UPosition int
	Position  string
}

type DeviceModel struct {
	sync.RWMutex
	DeviceModelId   int
	DeviceModelName string
	DeviceType      string
	U         		uint8
	HALF_FULL		string
}

type Cabinet struct {
	sync.RWMutex
	CabinetId     int
	CabinetName   string
	CapacityTotal uint32
	CapacityUsed  uint32
	Idc     	  *Idc
}
type Idc struct{
	sync.RWMutex
	IdcId int
	IdcName string
	Location *Location
}

type Location struct {
	sync.RWMutex
	LocationId         int
	LocationName       string
	FatherLocationId   interface{}
}


type Tag string

type Prometheus struct {
	dbobj db.Dbi
	ReadCache bool
}

var PROMETHEUS 				*Prometheus

func Init(mainCfg cfg.MainCfg) {
	var err error
	log.Debug("Start init Database.")
	PROMETHEUS = new(Prometheus)
	PROMETHEUS.dbobj, err = db.Init(mainCfg.DbCfg)
	PROMETHEUS.ReadCache=mainCfg.PrometheusCfg.ReadCache
	if err != nil {
		log.Fatal("Init Database Failed.")
		os.Exit(255)
	}
	log.Debug("Success init Database.")
	log.Debug("Start Open Database.")
	err = PROMETHEUS.dbobj.Start()
	if err != nil {
		log.Error("Open Database failed=>", err.Error())
		os.Exit(255)
	}
	log.Debug("Open Database Success")
	log.Debug("Init Data Start")
	err=Cache_Index_Init()
	if err!=nil{
		log.Fatal("Init Data Failed")
		os.Exit(255)
	}
	log.Debug("Init Data Success")
}


//func if_device_name_exist(name string)bool{
//	for _,a_kind_of_device :=range DEVICE_INDEX_NAME{
//		_,exist:=a_kind_of_device[name]
//		if exist{
//			return true
//		}
//	}
//	return false
//}

func ifcache()bool{
	return PROMETHEUS.ReadCache
}
//func (device *Device)Init(	deviceId  int,deviceName   string,deviceType  string,fatherDeviceId interface{}){
//	device.DeviceId=deviceId
//	device.DeviceName=deviceName
//	device.DeviceType=deviceType
//	device.FatherDeviceId=fatherDeviceId
//}

