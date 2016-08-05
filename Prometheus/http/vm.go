package http

import (
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

func GetVm(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == ""{
		r, err = prometheus.GetVm(nil)
	} else {
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = prometheus.GetVm(nil,id_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id !="" {
		if len(r.([]*prometheus.Vm))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Vm)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddVm(ctx *macaron.Context) {
	deviceName:=ArgString{"DeviceName",true,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	fatherDeviceId:=ArgInt{"FatherDeviceId",true,nil}
	deviceModelId:=ArgInt{"DeviceModelId",true,nil}
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{fatherDeviceId,deviceModelId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	err=prometheus.AddVm(args_string["DeviceName"].(string),args_int["DeviceModelId"].(int),args_int["FatherDeviceId"].(int))
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201,"Add Success")
	}
}
