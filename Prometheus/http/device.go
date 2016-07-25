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
		r, err = prometheus.GetDevice(nil)
	} else {
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = prometheus.GetDevice(nil,id_int)
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
	deviceName:=ArgString{"DeviceName",true,nil}
	deviceType:=ArgString{"DeviceType",true,nil}
	faterDeviceId:=ArgInt{"FatherDeviceId",false,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName,deviceType})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{faterDeviceId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	device:=prometheus.Device{DeviceName:args_string["DeviceName"].(string),DeviceType:args_string["DeviceType"].(string),FatherDeviceId:args_int["FatherDeviceId"]}
	err=device.AddDevice()
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
	device:=&prometheus.Device{DeviceId:id_int}
	err = device.DeleteDevice()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateDevice(ctx *macaron.Context) {
	deviceId:= ctx.Params("id")
	deviceId_int, err := strconv.Atoi(deviceId)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	devices,err:=prometheus.GetDevice(nil,deviceId_int)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	if len(devices) == 0{
		ctx.JSON(404, "Not Found")
		return
	}
	device:=devices[0]
	deviceName:=ArgString{"DeviceName",false,device.DeviceName}
	deviceType:=ArgString{"DeviceType",false,device.DeviceType}
	faterDeviceId:=ArgInt{"FatherDeviceId",false,device.FatherDeviceId}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName,deviceType})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{faterDeviceId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	device.DeviceName=args_string["DeviceName"].(string)
	device.DeviceType=args_string["DeviceType"].(string)
	device.FatherDeviceId=args_int["FatherDeviceId"]
	err = device.UpdateDevice()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}
