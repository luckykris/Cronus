package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetIdc(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("id") == ""{
		r, err = prometheus.GetIdc(nil)
	} else {
		r, err = prometheus.GetIdc(ctx.ParamsInt("id"))
		if len(r.([]*prometheus.Idc))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Idc)[0]
		}
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}

