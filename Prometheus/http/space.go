package http

import (
	//"fmt"
	"strconv"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetSpace(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	cabinet_id := ctx.Req.Form.Get("cabinet_id")
	device_id := ctx.Req.Form.Get("device_id")
	var r interface{}
	var err error
	var condition = []string{}
	if cabinet_id != "" {
		condition = append(condition, "cabinet_id = "+cabinet_id)
	}
	if device_id != "" {
		condition = append(condition, "device_id = "+device_id)
	}
	r, err = prometheus.GetSpace(condition...)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}



func GetDeviceSpace(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	r, err = device.GetSpace()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}