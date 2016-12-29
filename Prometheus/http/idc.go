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
	if ctx.Params("*") == ""{
		r, err = prometheus.GetIdc(nil,nil)
	} else {
		r, err = prometheus.GetOneIdc(nil,ctx.Params("*"))
		if err!=nil{
			ctx.JSON(404, err.Error())
			return
		}
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}

func AddIdc(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var err error
	self:=new(prometheus.Idc)
	IdcName     ,err := getArg(ctx,"IdcName"        ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    LocationName,err := getArg(ctx,"LocationName"   ,STRING,false,nil);if err != nil {ctx.JSON(405, err.Error());return}
    if err != nil {ctx.JSON(405, err.Error());return}
    Location, err := prometheus.GetOneLocation(nil,LocationName)
	self.Location=Location
	self.IdcName =IdcName.(string)
	err=prometheus.AddIdc(self)
	if err != nil {ctx.JSON(500, err.Error());return}
	ctx.JSON(200, SUCCESS)
}

func UpdateIdc(ctx *macaron.Context) {
	var err error
	var self,fake *prometheus.Idc
	self, err = prometheus.GetOneIdc(nil,ctx.Params("*"))
	if err!=nil{
		ctx.JSON(404, err.Error())
		return
	}
	fake=self.FakeCopy()
	ctx.Req.ParseForm()
    IdcName,err := getArg(ctx,"IdcName" ,STRING,false,fake.IdcName);if err != nil {ctx.JSON(405, err.Error());return}
    fake.IdcName=IdcName.(string)
    err=self.Update(fake)	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}

func DeleteIdc(ctx *macaron.Context) {
	var err error
	var self  *prometheus.Idc
	self, err = prometheus.GetOneIdc(nil,ctx.Params("*"))
	if err!=nil   {ctx.JSON(404, err.Error());return}
    err=self.Delete()	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}