package http

import (
	//"fmt"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)


func GetCabinetSpace(ctx *macaron.Context) {
	var r interface{}
	var err error
	var cabinet *prometheus.Cabinet
	cabinets, err := prometheus.GetCabinet(ctx.ParamsInt("id"),nil)
	if len(cabinets)<1 {
		ctx.JSON(404, err.Error())
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

func GetCabinet(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("*") == ""{
		r, err = prometheus.GetCabinet(nil,nil)
	} else {
		r, err = prometheus.GetOneCabinet(nil,ctx.Params("*"))
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

func AddCabinet(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var err error
	self:=new(prometheus.Cabinet)
	CabinetName  ,err := getArg(ctx,"CabinetName"   ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    CapacityTotal ,err := getArg(ctx,"CapacityTotal" ,INT   ,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    IdcName       ,err := getArg(ctx,"IdcName"       ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    if err != nil {ctx.JSON(405, err.Error());return}
    Idc,err:=prometheus.GetOneIdc(nil,IdcName)
    if err != nil {ctx.JSON(404, err.Error());return}	
	self.CabinetName=CabinetName.(string)
	self.CapacityTotal=uint32(CapacityTotal.(int))
	self.Idc=Idc
	err=prometheus.AddCabinet(self)
	if err != nil {ctx.JSON(500, err.Error());return}
	ctx.JSON(200, SUCCESS)
}

func UpdateCabinet(ctx *macaron.Context) {
	var err error
	var self,fake *prometheus.Cabinet
	self, err = prometheus.GetOneCabinet(nil,ctx.Params("*"))
	if err!=nil{
		ctx.JSON(404, err.Error())
		return
	}
	fake=self.FakeCopy()
	ctx.Req.ParseForm()
    CabinetName,err := getArg(ctx,"CabinetName" ,STRING,false,fake.CabinetName);if err != nil {ctx.JSON(405, err.Error());return}
    fake.CabinetName=CabinetName.(string)
    err=self.Update(fake)	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}

func DeleteCabinet(ctx *macaron.Context) {
	var err error
	var self  *prometheus.Cabinet
	self, err = prometheus.GetOneCabinet(nil,ctx.Params("*"))
	if err!=nil{ctx.JSON(404, err.Error());return}
    err=self.Delete()	
    if err != nil {ctx.JSON(500, err.Error());return}
    ctx.JSON(200,SUCCESS)
}