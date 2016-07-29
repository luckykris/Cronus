package http

import (
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

func GetServer(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == ""{
		r, err = prometheus.GetServer(nil)
	} else {
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = prometheus.GetServer(nil,id_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id !="" {
		if len(r.([]*prometheus.Server))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Server)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddServer(ctx *macaron.Context) {
	deviceName:=ArgString{"DeviceName",true,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	server:=prometheus.NewServer(args_string["DeviceName"].(string))
	err=server.AddServer()
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201,"Add Success")
	}
}
func DeleteServer(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	device:=&prometheus.Device{DeviceId:id_int}
	err = device.DeleteDevice()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}
