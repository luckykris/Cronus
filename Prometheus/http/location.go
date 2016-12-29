package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)

//location
func GetLocation(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("*") == ""{
		r, err = prometheus.GetLocation(nil,nil)
	} else {
		r, err = prometheus.GetOneLocation(nil,ctx.Params("*"))
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

func AddLocation(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var err error
	self:=new(prometheus.Location)
	LocationName     ,err := getArg(ctx,"LocationName"       ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    FatherLocationId ,err := getArg(ctx,"FatherLocationId"   ,INT   ,false,nil);if err != nil {ctx.JSON(405, err.Error());return}
    if err != nil {ctx.JSON(405, err.Error());return}
	self.LocationName=LocationName.(string)
	self.FatherLocationId=FatherLocationId
	err=prometheus.AddLocation(self)
	if err != nil {ctx.JSON(500, err.Error());return}
	ctx.JSON(200, SUCCESS)
}

func UpdateLocation(ctx *macaron.Context) {
	var err error
	var self,fake *prometheus.Location
	self, err = prometheus.GetOneLocation(nil,ctx.Params("*"))
	if err!=nil{
		ctx.JSON(404, err.Error())
		return
	}
	fake=self.FakeCopy()
	ctx.Req.ParseForm()
    LocationName,err := getArg(ctx,"LocationName" ,STRING,false,fake.LocationName);if err != nil {ctx.JSON(405, err.Error());return}
    fake.LocationName=LocationName.(string)
    err=self.Update(fake)	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}

func DeleteLocation(ctx *macaron.Context) {
	var err error
	var self  *prometheus.Location
	self, err = prometheus.GetOneLocation(nil,ctx.Params("*"))
	if err!=nil{ctx.JSON(404, err.Error());return}
    err=self.Delete()	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}