package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetCabinet(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("id") == ""{
		r, err = prometheus.GetCabinet(nil)
	} else {
		r, err = prometheus.GetCabinet(ctx.ParamsInt("id"))
		if len(r.([]*prometheus.Cabinet))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Cabinet)[0]
		}
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}


func GetCabinetSpace(ctx *macaron.Context) {
	var r interface{}
	var err error
	var cabinet *prometheus.Cabinet
	cabinets, err := prometheus.GetCabinet(ctx.ParamsInt("id"))
	if len(cabinets)<1 {
		ctx.JSON(404, "Not Found")
		return
	}else{
		cabinet = cabinets[0]
	}
	r,err=cabinet.GetSpace()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}