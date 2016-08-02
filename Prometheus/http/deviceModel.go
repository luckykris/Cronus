package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//Device Model
func GetDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetDeviceModel()
	} else {
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(404, "Not Found")
			return
		}
		r, err = prometheus.GetDeviceModel(id_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id != "" {
		if len(r.([]*prometheus.DeviceModel)) < 1 {
			ctx.JSON(404, "Not Found")
			return
		} else {
			r = r.([]*prometheus.DeviceModel)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddDeviceModel(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("DeviceModelName")
	_type := ctx.Req.Form.Get("DeviceType")
	u := ctx.Req.Form.Get("U")
	u_int, err := strconv.Atoi(u)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err= prometheus.AddDeviceModel(name, _type,u_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteDeviceModel(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateDeviceModel(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("DeviceModelName")
	_type := ctx.Req.Form.Get("DeviceType")
	err = _CheckHasClounms("device_model_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("device_type", _type, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateDeviceModel(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}
