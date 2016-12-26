package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

func GetDeviceModel(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("*") == ""{
		r, err = prometheus.GetDeviceModel()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
	} else {
		r, err = prometheus.GetOneDeviceModel(nil,ctx.Params("*"))
		if err!=nil{
			ctx.JSON(404, err.Error())
			return
		}
	}
	ctx.JSON(200, &r)
}

func AddDeviceModel(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var r interface{}
	var err error
	self:=new(prometheus.DeviceModel)
	DeviceModelName,err := getArg(ctx,"DeviceModelName" ,STRING,true ,nil    );if err != nil {ctx.JSON(405, err.Error());return}
	DeviceType     ,err := getArg(ctx,"DeviceType"      ,STRING,true ,nil    );if err != nil {ctx.JSON(405, err.Error());return}
    U              ,err := getArg(ctx,"U"               ,INT   ,false,0      );if err != nil {ctx.JSON(405, err.Error());return}
    HALF_FULL      ,err := getArg(ctx,"HALF_FULL"       ,STRING,false,"full" );if err != nil {ctx.JSON(405, err.Error());return}
    self.DeviceModelName=DeviceModelName.(string)
    self.DeviceType     =DeviceType.(string)
    self.U              =uint8(U.(int))
    self.HALF_FULL      =HALF_FULL.(string)
	err=prometheus.AddDeviceModel(self)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}

func DeleteDeviceModel(ctx *macaron.Context) {
	self,err := prometheus.GetOneDeviceModel(nil,ctx.Params("*"))
	if err!=nil {
		ctx.JSON(404, err.Error())
		return
	}
	err = self.Delete()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, SUCCESS)
}