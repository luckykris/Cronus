package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

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
	deviceName:=ArgString{"deviceName",true}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201, args_string)
		//ctx.JSON(201, "Add Success")
	}
//	father_device_id,error := arg2IntOrNil(ctx.Req.Form.Get("father_device_id"))
//	var father_device_id_int interface{}
//	var err error
//	if father_device_id == "" || father_device_id == "null" {
//		father_device_id_int = nil
//	} else {
//		father_device_id_int, err = strconv.Atoi(father_device_id)
//		if err != nil {
//			ctx.JSON(400, err.Error())
//			return
//		}
//	}
//	device:=&prometheus.Device{DeviceName:device_name,FatherDevice}
//	err = prometheus.AddDevice([][]interface{}{[]interface{}{device_name, , father_device_id_int}})
//	if err != nil {
//		ctx.JSON(400, err.Error())
//	} else {
//		ctx.JSON(201, "Add Success")
//	}
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
	father_device_id := ctx.Req.Form.Get("FatherLocationId")
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
	err = _CheckHasClounms("father_location_id", father_device_id, true, &cloumns, &values)
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
