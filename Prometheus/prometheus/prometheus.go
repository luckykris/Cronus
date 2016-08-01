package prometheus

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	"os"
)

type Server struct {
	Serial string
	Hostname string
	Memsize int
	Os  string
	Release float64
	LastChangeTime int64
	Checksum string
	Device 
}

type Device struct {
	DeviceId       int
	DeviceName     string
	FatherDeviceId interface{}
	DeviceModel  *DeviceModel
	NetPorts	[]NetPort
}

type NetPort struct {
	NetPortId int
	Mac       interface{}
	Ipv4Int   interface{}
	Type      string
}

type Space struct {
	CabinetId int
	DeviceId  int
	UPosition int
	Position  string
}

type DeviceModel struct {
	DeviceModelId   int
	DeviceModelName string
	DeviceType      string
}

type Cabinet struct {
	CabinetId     int
	CabinetName   string
	IsCloud       string
	CapacityTotal uint64
	CapacityUsed  uint64
	LocationId    int
}

type Location struct {
	LocationId       int
	LocationName     string
	Picture          string
	FatherLocationId interface{}
}

type Tag struct {
	TagId int
	TagName   string
}

type Prometheus struct {
	dbobj db.Dbi
	DeviceModelMapId map[int]*DeviceModel
	ServerMapId map[int]*Server
	NetPortMap map[int]*NetPort
	DeviceNameMap map[string]*Device
	DeviceIdMap map[int]*Device
}

var PROMETHEUS *Prometheus

func Init(mainCfg cfg.MainCfg) {
	var err error
	log.Debug("Start init Database.")
	PROMETHEUS = &Prometheus{ServerMapId:map[int]*Server{},
							 DeviceNameMap:map[string]*Device{},
							 DeviceIdMap:map[int]*Device{},
							 DeviceModelMapId:map[int]*DeviceModel{},
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
	log.Debug("Load Server Start")
	err=LoadCache()
	if err!=nil{
		log.Fatal("Load Server Failed")
		os.Exit(255)
	}
	log.Debug("Load Server Success")
	fmt.Printf("%#v",PROMETHEUS.ServerMapId)
}


//func (device *Device)Init(	deviceId  int,deviceName   string,deviceType  string,fatherDeviceId interface{}){
//	device.DeviceId=deviceId
//	device.DeviceName=deviceName
//	device.DeviceType=deviceType
//	device.FatherDeviceId=fatherDeviceId
//}

