package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)


func GetTag(ctx *macaron.Context) {
	tag_id := ctx.Params("id")
	var r interface{}
	var err error
	if tag_id == ""{
		r, err = prometheus.GetTag()
	} else {
		tag_id_int, err := strconv.Atoi(tag_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = prometheus.GetTag(tag_id_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if tag_id !="" {
		if len(r.([]prometheus.Tag))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]prometheus.Tag)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddTag(ctx *macaron.Context) {
	tagName:=ArgString{"TagName",true,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{tagName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	tag:=prometheus.Tag{TagName:args_string["TagName"].(string)}
	err=prometheus.AddTag(&tag)
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201,"Add Success")
	}
}
func DeleteTag(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	tag:=&prometheus.Tag{TagId:id_int}
	err = tag.DeleteTag()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateTag(ctx *macaron.Context) {
	tag_id:= ctx.Params("id")
	tag_id_int, err := strconv.Atoi(tag_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	tags,err:=prometheus.GetTag(tag_id_int)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	if len(tags) == 0{
		ctx.JSON(404, "Not Found")
		return
	}
	tag:=tags[0]
	tagName:=ArgString{"TagName",false,tag.TagName}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{tagName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	tag.TagName=args_string["TagName"].(string)
	err = tag.UpdateTag()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}



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
	tag := &prometheus.Tag{TagId:tag_id_int}
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = device.AddTag(tag)
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200,"Success")
	}
}

func DeleteDeviceTag(ctx *macaron.Context) {
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
	tag := &prometheus.Tag{TagId:tag_id_int}
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = device.DeleteTag(tag)
	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200,"Success")
	}
}