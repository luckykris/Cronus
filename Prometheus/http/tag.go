package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)





func GetDeviceTag(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag_id := ctx.Params("TagId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	if tag_id != "" {
		tag_id_int, err := strconv.Atoi(tag_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = device.GetTag(tag_id_int)
	} else {
		r, err = device.GetTag()
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if tag_id != "" {
		if len(r.([]prometheus.Tag)) < 1 {
			ctx.JSON(404, "Not Found")
		} else {
			r = r.([]prometheus.Tag)[0]
			ctx.JSON(200, &r)
		}
	} else {
		ctx.JSON(200, &r)
	}
}
func AddDeviceTag(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag_id := ctx.Params("TagId")
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	tag_id_int, err := strconv.Atoi(tag_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = device.AddTag(tag_id_int)
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200,"Success")
	}
}
