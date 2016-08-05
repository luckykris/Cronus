package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)





func DeleteTag(ctx *macaron.Context) {
	tag := ctx.Params("tag")
	err := prometheus.DeleteTag(prometheus.Tag(tag))
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}





func GetDeviceTag(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device,exist := prometheus.PROMETHEUS.DeviceMapId[device_id_int]
	if !exist{
		ctx.JSON(404, "Device Not Found")
		return
	}
	r, err = device.GetTag()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200,r)
}
func AddDeviceTag(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag := ctx.Params("Tag")
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device,exist := prometheus.PROMETHEUS.DeviceMapId[device_id_int]
	if !exist{
		ctx.JSON(404, "Device Not Found")
		return
	}
	err = device.AddTag(prometheus.Tag(tag))
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200,"Success")
	}
}

func DeleteDeviceTag(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag := ctx.Params("Tag")
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device,exist := prometheus.PROMETHEUS.DeviceMapId[device_id_int]
	if !exist{
		ctx.JSON(404, "Device Not Found")
		return
	}
	err = device.DeleteTag(prometheus.Tag(tag))
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200,"Success")
	}
}