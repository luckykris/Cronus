package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetDevice(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetDevice()
	} else {
		r, err = prometheus.GetDevice("device_id=" + id)

	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id !="" {
		if len(r.([]prometheus.Device))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]prometheus.Device)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddDevice(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("LocationName")
	pic := ctx.Req.Form.Get("Picture")
	father_id := ctx.Req.Form.Get("FatherLocationId")
	var father_id_int interface{}
	var err error
	if father_id == "" || father_id == "null" {
		father_id_int = nil
	} else {
		father_id_int, err = strconv.Atoi(father_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
	}
	err = prometheus.AddLocation([][]interface{}{[]interface{}{name, pic, father_id_int}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteDevice(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteLocation(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateDevice(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("LocationName")
	pic := ctx.Req.Form.Get("Picture")
	father_id := ctx.Req.Form.Get("FatherLocationId")
	err = _CheckHasClounms("location_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("picture", pic, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("father_location_id", father_id, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateLocation(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}
