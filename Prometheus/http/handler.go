package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

func GetDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetDeviceModel()
	} else {
		r, err = prometheus.GetDeviceModel([]byte("device_model_id=" + id))
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	switch len(r.([]global.DeviceModel)) {
	case 0:
		ctx.JSON(404, "Not Found")
	case 1:
		r = r.([]global.DeviceModel)[0]
		ctx.JSON(200, &r)
	default:
		ctx.JSON(200, &r)
	}
}

func NotFound(ctx *macaron.Context) {
	ctx.JSON(404, "No Such Url")
}
