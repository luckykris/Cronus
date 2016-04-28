package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	netPort_id := ctx.Params("NetPortId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	if netPort_id != "" {
		netPort_id_int, err := strconv.Atoi(netPort_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = device.GetNetPort(netPort_id_int)
	} else {
		r, err = device.GetNetPort()
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if netPort_id != "" {
		if len(r.([]prometheus.NetPort)) < 1 {
			ctx.JSON(404, "Not Found")
		} else {
			r = r.([]prometheus.NetPort)[0]
			ctx.JSON(200, &r)
		}
	} else {
		ctx.JSON(200, &r)
	}
}
