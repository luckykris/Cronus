package http

import (
	//"fmt"
	"path/filepath"
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)



func GetNetPort(ctx *macaron.Context) {
	if filepath.HasPrefix(ctx.Req.Request.RequestURI,server_path){
		self, err := prometheus.GetOneServer(nil,ctx.Params("*"))
		if err!=nil {
			ctx.JSON(404, err.Error())
		}else{
			ctx.JSON(200, self.GetNetPort())
		}
		return
	}else{
		ctx.JSON(503, "url path parse fatal")
		return
	}
}

//func AddNetPort(ctx *macaron.Context) {
//	ctx.Req.ParseForm()
//	var err error
//	server:=new(prometheus.Server)
//	DeviceName     ,err := getArg(ctx,"Ipv4"             ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
//    DeviceModelId  ,err := getArg(ctx,"NetPortType"      ,STRING,false,"ether");if err != nil {ctx.JSON(405, err.Error());return}
//   
//	server.Device.DeviceName=DeviceName.(string)
//	server.Device.DeviceModel,err=prometheus.GetOneDeviceModel(DeviceName,DeviceModelId)
//	if err != nil {
//		ctx.JSON(405, err.Error())
//		return
//	}
//	server.Device.GroupId=0
//	server.Device.Env=0
//	err=prometheus.AddServer(server)
//	if err != nil {
//		ctx.JSON(500, err.Error())
//		return
//	}
//	ctx.JSON(200, SUCCESS)
//}

//func DeleteServer(ctx *macaron.Context) {
//	var err error
//	var self *prometheus.Server
//	self, err = prometheus.GetOneServer(nil,ctx.Params("*"))
//	if err!=nil {
//		ctx.JSON(404, "Not Found")
//		return
//	}
//	err = self.Delete()
//	if err != nil {
//		ctx.JSON(500, err.Error())
//		return
//	}
//	ctx.JSON(200, nil)
//}
//
//func UpdateServer(ctx *macaron.Context) {
//	var err error=nil
//	var server,fake *prometheus.Server
//	server, err = prometheus.GetOneServer(nil,ctx.Params("*"))
//	if err!=nil{
//		ctx.JSON(404, err.Error())
//		return
//	}
//	fake=server.FakeCopy()
//	ctx.Req.ParseForm()
//    DeviceName,err := getArg(ctx,"DeviceName" ,STRING,false,fake.Get_DeviceName());if err != nil {ctx.JSON(405, err.Error());return}
//    GroupId   ,err := getArg(ctx,"GroupId"    ,INT   ,false,fake.Get_GroupId())   ;if err != nil {ctx.JSON(405, err.Error());return}
//    Env       ,err := getArg(ctx,"Env"        ,INT   ,false,fake.Get_Env())	      ;if err != nil {ctx.JSON(405, err.Error());return}
//    Serial    ,err := getArg(ctx,"Serial" ,STRING,false,fake.Serial)              ;if err != nil {ctx.JSON(405, err.Error());return}
//	Hostname  ,err := getArg(ctx,"Hostname" ,STRING,false,fake.Hostname)          ;if err != nil {ctx.JSON(405, err.Error());return}
//	Memsize   ,err := getArg(ctx,"Memsize" ,INT,false,fake.Memsize)               ;if err != nil {ctx.JSON(405, err.Error());return}
//	Os        ,err := getArg(ctx,"Os" ,STRING,false,fake.Os)                      ;if err != nil {ctx.JSON(405, err.Error());return}
//	Release   ,err := getArg(ctx,"Release" ,FLOAT,false,fake.Release)             ;if err != nil {ctx.JSON(405, err.Error());return}
//
//	fake.Device.DeviceName=DeviceName.(string)
//	fake.Device.GroupId   =GroupId.(int)
//	fake.Device.Env       =uint8(Env.(int))
//	fake.Serial    =Serial.(string)
//	fake.Hostname  =Hostname.(string)
//	fake.Memsize   =uint32(Memsize.(int))
//	fake.Os        =Os.(string)
//	fake.Release   =Release.(float64)
//	err=server.Update(fake)
//	if err != nil {
//		ctx.JSON(500, err.Error())
//		return
//	}
//	ctx.JSON(200, err)
//}
