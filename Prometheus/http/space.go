package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetSpace(ctx *macaron.Context) {
	id := ctx.Params("id")
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("cabinet_id")
	pic := ctx.Req.Form.Get("device_id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetSpace()
	} else {
		r, err = prometheus.GetSpace("device_id=" + id)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id !="" {
		if len(r.([]global.Device))<1 {
			ctx.JSON(404, "Not Found")
		}else{
			r = r.([]global.Device)[0]
		}
	}
	ctx.JSON(200, &r)
}

