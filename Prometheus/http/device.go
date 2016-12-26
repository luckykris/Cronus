package http

import (
	"path/filepath"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)



func DeleteDevice(ctx *macaron.Context) {
	if filepath.HasPrefix(ctx.Req.Request.RequestURI,server_path){
		self, err := prometheus.GetOneServer(nil,ctx.Params("*"))
		if err!=nil {
			ctx.JSON(404, err.Error())
		}else{
			err = self.Delete()
			if err!=nil{
				ctx.JSON(500,err.Error())
			}else{	
				ctx.JSON(200,SUCCESS)
			}
		}
	}else{
		ctx.JSON(503,"url path parse fatal")
	}
}
