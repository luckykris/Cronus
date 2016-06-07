package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	netPort_id := ctx.Params("NetPortId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	if netPort_id != "" {
		netPort_id_int, err := strconv.Atoi(netPort_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		r, err = device.GetNetPort(netPort_id_int)
	} else {
		r, err = device.GetNetPort()
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if netPort_id != "" {
		if len(r.([]prometheus.NetPort)) < 1 {
			ctx.JSON(404, "Not Found")
		} else {
			r = r.([]prometheus.NetPort)[0]
			ctx.JSON(200, &r)
		}
	} else {
		ctx.JSON(200, &r)
	}
}

func AddNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	mac:=ArgString{"Mac",false,nil}
	ipv4Int:=ArgInt{"Ipv4Int",false,nil}
	_type:=ArgString{"Type",true,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{mac,_type})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{ipv4Int})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	netPort:=prometheus.NetPort{Mac:args_string["Mac"],  Ipv4Int:args_int["Ipv4Int"],  Type:args_string["Type"].(string)}	
	device:=&prometheus.Device{DeviceId:device_id_int}
	err = device.AddNetPort(netPort)
	if err!=nil{
		ctx.JSON(400, err.Error())
	}else{
		ctx.JSON(201,"Add Success")
	}
}

func UpdateNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	netPort_id:= ctx.Params("NetPortId")
	device_id_int, err := strconv.Atoi(device_id)
	netPort_id_int, err := strconv.Atoi(netPort_id)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	device:=&prometheus.Device{DeviceId:device_id_int}
	netPorts,err:=device.GetNetPort(netPort_id_int)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	if len(netPorts) == 0{
		ctx.JSON(404, "Not Found")
		return
	}
	netPort:=netPorts[0]
	mac:=ArgString{"Mac",false,netPort.Mac}
	ipv4Int:=ArgInt{"Ipv4Int",false,netPort.Ipv4Int}
	_type:=ArgString{"Type",false,netPort.Type}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{mac,_type})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}	
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{ipv4Int})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	netPort.Mac=args_string["Mac"]
	netPort.Ipv4Int=args_int["Ipv4Int"]
	netPort.Type=args_string["Type"].(string)
	err = netPort.UpdateNetPort()
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Update Success")
	}
}

func DeleteNetPort(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	netPort_id:= ctx.Params("NetPortId")
	device_id_int, err := strconv.Atoi(device_id)
	netPort_id_int, err := strconv.Atoi(netPort_id)
	if err!=nil{
		ctx.JSON(500, err.Error())
		return
	}
	device:=&prometheus.Device{DeviceId:device_id_int}
	netPort:=&prometheus.NetPort{NetPortId:netPort_id_int}
	err = device.DeleteNetPort(netPort)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}