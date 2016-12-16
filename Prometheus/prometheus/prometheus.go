package prometheus

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	"os"
	"sync"
)
type Device_i interface{
	GetTag() ([]Tag, error)
	AddTag(...Tag) error
	DeleteTag(...Tag) error
	ComputSum()string
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
	Memsize int
	Os  string
	Release float64
	LastChangeTime int64
	Checksum string
	Device 
}

type NetPort struct {
	Mac       interface{}
	Ipv4Int   interface{}
	NetPortType      string
}

type Space struct {
	CabinetId int
	DeviceId  int
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
	CabinetId     int
	CabinetName   string
	IsCloud       string
	CapacityTotal uint64
	CapacityUsed  uint64
	IdcId    int
}
type Idc struct{
	IdcId int
	IdcName string
	LocationId int
}

type Location struct {
	LocationId       int
	LocationName     string
	Picture          string
	FatherLocationId interface{}
}

type Tag string

type Prometheus struct {
	dbobj db.Dbi
	DeviceModelMapId map[int]*DeviceModel
	DeviceCache map[string]map[int]Device_i
	CabinetMapId map[int]*Cabinet
	LocationMapId map[int]*Location
	IdcMapId map[int]*Idc
}

var PROMETHEUS 				*Prometheus

func Init(mainCfg cfg.MainCfg) {
	var err error
	log.Debug("Start init Database.")
	PROMETHEUS = &Prometheus{DeviceCache:map[string]map[int]Device_i{"server":map[int]Device_i{},"vm":map[int]Device_i{}},
							 DeviceModelMapId:map[int]*DeviceModel{},
							 CabinetMapId:map[int]*Cabinet{},
							 LocationMapId:map[int]*Location{},
							 IdcMapId:map[int]*Idc{},
							}
	PROMETHEUS.dbobj, err = db.Init(mainCfg.DbCfg)
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
	fmt.Printf("%#v",PROMETHEUS.DeviceCache)
}


func if_device_name_exist(name string)bool{
	for _,a_kind_of_device :=range DEVICE_INDEX_NAME{
		_,exist:=a_kind_of_device[name]
		if exist{
			return true
		}
	}
	return false
}

//func (device *Device)Init(	deviceId  int,deviceName   string,deviceType  string,fatherDeviceId interface{}){
//	device.DeviceId=deviceId
//	device.DeviceName=deviceName
//	device.DeviceType=deviceType
//	device.FatherDeviceId=fatherDeviceId
//}

