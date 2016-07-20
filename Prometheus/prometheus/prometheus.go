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
	Memsize uint64
	Os  string
	Release float32
	LastChangeTime uint64
	Checksum string
	Device
}

type Device struct {
	DeviceId       int
	DeviceName     string
	DeviceType  string
	FatherDeviceId interface{}
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
}

var PROMETHEUS *Prometheus

func Init(mainCfg cfg.MainCfg) {
	var err error
	log.Debug("Start init Database.")
	PROMETHEUS = &Prometheus{}
	PROMETHEUS.dbobj, err = db.Init(mainCfg.DbCfg)
	if err != nil {
		fmt.Println(err)
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
}
