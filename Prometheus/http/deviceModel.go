package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetDeviceModel(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("id") == ""{
		r, err = prometheus.GetDeviceModel()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
	} else {
		r, err = prometheus.GetOneDeviceModel(ctx.ParamsInt("id"))
		if err!=nil{
			ctx.JSON(404, err.Error())
			return
		}
	}
	ctx.JSON(200, &r)
}
