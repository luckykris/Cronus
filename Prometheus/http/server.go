package http

import (
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

func GetServer(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("id") == ""{
		r, err = prometheus.GetServer(nil,nil)
	} else {
		r, err = prometheus.GetServer(nil,ctx.ParamsInt("id"))
		if len(r.([]*prometheus.Server))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Server)[0]
		}
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}

func AddServer(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var r interface{}
	var err error
	server:=new(prometheus.Server)
	server.Device.DeviceName=ctx.Query("DeviceName")
	server.Device.DeviceModel,err=prometheus.GetDeviceModel(ctx.QueryInt("DeviceModelId"))
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	server.Device.GroupId=0
	server.Device.Env=0
	err=prometheus.AddServer(server)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}
