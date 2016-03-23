package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetAllDeviceType() ([]global.DeviceType, error) {
	r, err := PROMETHEUS.dbobj.GetDeviceType()
	return r, err
}
