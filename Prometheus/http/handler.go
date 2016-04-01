package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
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
	name := ctx.Req.Form.Get("Name")
	_type := ctx.Req.Form.Get("Type")
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
	name := ctx.Req.Form.Get("Name")
	_type := ctx.Req.Form.Get("Type")
	_CheckHasClounms("device_model_name", name, &cloumns, &values)
	_CheckHasClounms("device_type", _type, &cloumns, &values)
	err = prometheus.UpdateDeviceModel(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

//cabinet
func GetCabinet(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetCabinet()
	} else {
		r, err = prometheus.GetCabinet("cabinet_id=" + id)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	switch len(r.([]global.Cabinet)) {
	case 0:
		ctx.JSON(404, "Not Found")
	case 1:
		r = r.([]global.Cabinet)[0]
		ctx.JSON(200, &r)
	default:
		ctx.JSON(200, &r)
	}
}
func AddCabinet(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("Name")
	iscloud := ctx.Req.Form.Get("IsCloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_total_int, err := strconv.Atoi(capacity_total)
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	capacity_used_int, err := strconv.Atoi(capacity_used)
	location_id := ctx.Req.Form.Get("LocationId")
	location_id_int, err := strconv.Atoi(location_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.AddCabinet([][]interface{}{[]interface{}{name, iscloud, capacity_total_int, capacity_used_int, location_id_int}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteCabinet(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteCabinet(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateCabinet(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("Name")
	iscloud := ctx.Req.Form.Get("IsVloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_total_int, err := strconv.Atoi(capacity_total)
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	capacity_used_int, err := strconv.Atoi(capacity_used)
	location_id := ctx.Req.Form.Get("LocationId")
	location_id_int, err := strconv.Atoi(location_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	_CheckHasClounms("cabinet_name", name, &cloumns, &values)
	_CheckHasClounms("iscloud", iscloud, &cloumns, &values)
	_CheckHasClounms("capacity_total", capacity_total_int, &cloumns, &values)
	_CheckHasClounms("capacity_used", capacity_used_int, &cloumns, &values)
	_CheckHasClounms("location_id", location_id_int, &cloumns, &values)
	err = prometheus.UpdateCabinet(id_int, cloumns, values)
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
