package http

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

func GetDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetDeviceModel()
	} else {
		r, err = prometheus.GetDeviceModel("device_model_id=" + id)
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

func AddDeviceModel(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("name")
	_type := ctx.Req.Form.Get("type")
	err := prometheus.AddDeviceModel([][]interface{}{[]interface{}{name, _type}})
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
	name := ctx.Req.Form.Get("name")
	fmt.Println(name)
	_type := ctx.Req.Form.Get("type")
	_CheckHasClounms("device_model_name", name, &cloumns, &values)
	_CheckHasClounms("device_type", _type, &cloumns, &values)
	fmt.Println(cloumns)
	err = prometheus.UpdateDeviceModel(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func _CheckHasClounms(cn string, v interface{}, cls *[]string, vals *[]interface{}) {
	switch v.(type) {
	case string:
		if v != "" {
			*cls = append(*cls, cn)
			*vals = append(*vals, v)
		}
	case int:
		*cls = append(*cls, cn)
		*vals = append(*vals, v)
	}
}

func NotFound(ctx *macaron.Context) {
	ctx.JSON(404, "No Such Url")
}
