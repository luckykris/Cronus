package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//location
func GetIdc(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetIdc(nil)
	} else {
		id_int, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(404, "Not Found")
			return
		}
		r, err = prometheus.GetIdc(nil,id_int)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	if id != "" {
		if len(r.([]prometheus.Location)) < 1 {
			ctx.JSON(404, "Not Found")
			return
		} else {
			r = r.([]prometheus.Location)[0]
		}
	}
	ctx.JSON(200, &r)
}