package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

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
	if id != "" {
		if len(r.([]global.Cabinet)) < 1 {
			ctx.JSON(404, "Not Found")
		} else {
			r = r.([]global.Cabinet)[0]
		}
	}
	ctx.JSON(200, &r)
}
func AddCabinet(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("CabinetName")
	iscloud := ctx.Req.Form.Get("IsCloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_total_int, err := strconv.Atoi(capacity_total)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	capacity_used_int, err := strconv.Atoi(capacity_used)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
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
	name := ctx.Req.Form.Get("CabinetName")
	iscloud := ctx.Req.Form.Get("IsCloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	location_id := ctx.Req.Form.Get("LocationId")
	err = _CheckHasClounms("cabinet_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("iscloud", iscloud, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("capacity_total", capacity_total, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("capacity_used", capacity_used, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("location_id", location_id, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateCabinet(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}
