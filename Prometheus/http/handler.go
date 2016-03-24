package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)


func  GetDevice(ctx *macaron.Context) {
	id:=ctx.Params("id")
	var r interface{}
	if id == ""{
		r,_=prometheus.GetAllDeviceType()
	}else{
		r=global.DeviceType{1,"test"}
	}
	ctx.JSON(200,&r)
}

//func NotFound(r macaron.Render) {
//	r.HTML(404, "notfound", nil)
//}
