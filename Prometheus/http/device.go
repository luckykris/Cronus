package http

import (
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

func GetDevice(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == ""{
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
	deviceName:=ArgString{"deviceName",true,nil}
	deviceModelId:=ArgInt{"deviceModelId",true,nil}
	faterDeviceId:=ArgInt{"fatherDeviceId",false,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{deviceModelId,faterDeviceId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	device:=prometheus.Device{DeviceName:args_string["deviceName"].(string),DeviceModelId:args_int["deviceModelId"].(int),FatherDeviceId:args_int["fatherDeviceId"]}
	err=prometheus.AddDevice(&device)
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201,"Add Success")
	}
}
func DeleteDevice(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteDevice(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateDevice(ctx *macaron.Context) {
	deviceId:= ctx.Params("id")
	devices,err:=prometheus.GetDevice("device_id="+deviceId)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	if len(devices) == 0{
		ctx.JSON(404, "Not Found")
		return
	}
	device:=devices[0]
	deviceName:=ArgString{"deviceName",false,device.DeviceName}
	deviceModelId:=ArgInt{"deviceModelId",false,device.DeviceModelId}
	faterDeviceId:=ArgInt{"fatherDeviceId",false,device.FatherDeviceId}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{deviceModelId,faterDeviceId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	device.DeviceName=args_string["deviceName"].(string)
	device.DeviceModelId=args_int["DeviceModelId"].(int)
	device.FatherDeviceId=args_int["fatherDeviceId"]
	err = device.UpdateDevice()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}
