package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	ipv4_int := ctx.Params("Ipv4Int")
	var r interface{}
	var err error
	if ipv4_int == "" {
		r, err = prometheus.GetNetPort("device_id=" + device_id)
	} else {
		r, err = prometheus.GetNetPort("device_id=" + device_id,"ipv4_int=" + ipv4_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if len(r.([]global.NetPort))<1 {
		ctx.JSON(404, "Not Found")
		return
	}else{
		r = r.([]global.NetPort)[0]
	}
	ctx.JSON(200, &r)
}

func AddNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag_id := ctx.Params("TagId")
	var err error
	err = prometheus.AddDeviceTag([][]interface{}{[]interface{}{device_id,tag_id}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	tag_id := ctx.Params("TagId")
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	tag_id_int, err := strconv.Atoi(tag_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}	
	err = prometheus.DeleteDeviceTag(device_id_int,tag_id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success") 
	}
}

