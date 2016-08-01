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
		if len(r.([]*prometheus.Device))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Device)[0]
		}
	}
	ctx.JSON(200, &r)
}

func AddDevice(ctx *macaron.Context) {
	deviceName:=ArgString{"DeviceName",true,nil}
	deviceModelId:=ArgInt{"DeviceModelId",true,nil}
	faterDeviceId:=ArgInt{"FatherDeviceId",false,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{faterDeviceId,deviceModelId})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	err=prometheus.AddDevice(args_string["DeviceName"].(string),args_int["DeviceModelId"].(int),args_int["FatherDeviceId"])
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

//func UpdateDevice(ctx *macaron.Context) {
//	deviceId:= ctx.Params("id")
//	deviceId_int, err := strconv.Atoi(deviceId)
//	if err != nil {
//		ctx.JSON(400, err.Error())
//		return
//	}
//	devices,err:=prometheus.GetDevice(nil,deviceId_int)
//	if err!=nil{
//		ctx.JSON(500, err.Error())
//		return
//	}
//	if len(devices) == 0{
//		ctx.JSON(404, "Not Found")
//		return
//	}
//	device:=devices[0]
//	deviceName:=ArgString{"DeviceName",false,device.DeviceName}
//	deviceModelId:=ArgString{"DeviceModelId",false,device.DeviceModelId}
//	faterDeviceId:=ArgInt{"FatherDeviceId",false,device.FatherDeviceId}
//	args_string,err:=getAllStringArgs(ctx,[]ArgString{deviceName})
//	if err!=nil{
//		ctx.JSON(400, err.Error())
//		return
//	}	
//	args_int,err:=getAllIntArgs(ctx,[]ArgInt{faterDeviceId,deviceModelId})
//	if err!=nil{
//		ctx.JSON(400, err.Error())
//		return
//	}	
//	device.DeviceName=args_string["DeviceName"].(string)
//	device.DeviceType=args_string["DeviceType"].(string)
//	device.FatherDeviceId=args_int["FatherDeviceId"]
//	err = device.UpdateDevice()
//	if err != nil {
//		ctx.JSON(400, err.Error())
//	} else {
//		ctx.JSON(204, "Update Success")
//	}
//}
