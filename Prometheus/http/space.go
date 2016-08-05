package http

import (
	//"fmt"
	"strconv"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetSpace(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	cabinet_id := ctx.Req.Form.Get("cabinet_id")
	device_id := ctx.Req.Form.Get("device_id")
	var r interface{}
	var err error
	var condition = []string{}
	if cabinet_id != "" {
		condition = append(condition, "cabinet_id = "+cabinet_id)
	}
	if device_id != "" {
		condition = append(condition, "device_id = "+device_id)
	}
	r, err = prometheus.GetSpace(condition...)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}



func GetDeviceSpace(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	var r interface{}
	var err error
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	device := &prometheus.Device{DeviceId: device_id_int}
	r, err = device.GetSpace()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}


func AddDeviceSpace(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	cabinet_id:=ArgInt{"CabinetId",true,nil}
	start_u:=ArgInt{"StartU",true,nil}
	args_int,err:=getAllIntArgs(ctx,[]ArgInt{cabinet_id,start_u})
	position:=ArgString{"Position",true,nil}
	args_string,err:=getAllStringArgs(ctx,[]ArgString{position})
	if err!=nil{
		ctx.JSON(400, err.Error())
		return
	}
	if device,ok:=prometheus.PROMETHEUS.ServerMapId[device_id_int];ok{
		err=device.AddSpace(args_int["CabinetId"].(int),args_int["StartU"].(int),args_string["Position"].(string))
		if err!=nil{
			ctx.JSON(400, err.Error())
			return
		}else{
			ctx.JSON(200, "Add Success")
			return
		}
	}	
}

func DeleteDeviceSpace(ctx *macaron.Context) {
	device_id := ctx.Params("DeviceId")
	device_id_int, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	if device,ok:=prometheus.PROMETHEUS.ServerMapId[device_id_int];ok{
		err=device.DeleteSpace()
		if err!=nil{
			ctx.JSON(400, err.Error())
			return
		}else{
			ctx.JSON(200, "Delete Success")
			return
		}
	}else{
		ctx.JSON(404, "Not Found")
		return
	}	
}



func GetCabinetSpace(ctx *macaron.Context) {
	cabinet_id := ctx.Params("CabinetId")
	var r interface{}
	var err error
	cabinet_id_int, err := strconv.Atoi(cabinet_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	cabinet,ok:=prometheus.PROMETHEUS.CabinetMapId[cabinet_id_int]
	if !ok{
		ctx.JSON(404, "Not Found")
		return
	}	
	r, err=cabinet.GetSpace()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}