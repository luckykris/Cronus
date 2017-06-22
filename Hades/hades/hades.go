package hades
import (
	log "github.com/Sirupsen/logrus"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/db"
	"os"
	"sync"
)
const (
	HALF = iota+1
	FULL
)
const (
	VM=iota
	SERVER
	CLOUD
	SWITCH
	ROUTER
)
type Container struct{}
type OS struct{
	id     int64
	hostname string
	memsize uint32
	processors []Processor
	portal []Portal
	os  string
	release float64
	lastChangeTime uint64
	checksum string
	boundDevice *Device
}
type Portal struct {
	mac       interface{}
	ipv4      uint32
	mask	  uint32
}
type Processor struct{}
type DeviceType struct{
	id   int64
	name string
	_type  uint8
	high	uint8
	width_f   uint8
	width_s   uint8
}
type Device struct{
	location *Location
	deviceType *DeviceType
}
type Location struct{
	start_u uint8
	rack *Rack
}
type Rack struct{
	idc *Idc
	row uint8
	high uint8
}







