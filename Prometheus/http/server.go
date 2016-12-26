package http

import (
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)


func GetServerSpace(ctx *macaron.Context) {
	var r interface{}
	var err error
	var server *prometheus.Server
	servers, err := prometheus.GetServer(nil,ctx.ParamsInt("*"))
	if len(servers)<1 {
		ctx.JSON(404, "Not Found")
		return
	}else{
		server = servers[0]
	}
	r,err=server.GetSpace()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, &r)
}
func GetServer(ctx *macaron.Context) {
	var r interface{}
	var err error
	if ctx.Params("*") == ""{
		r, err = prometheus.GetServer(nil,nil)
	} else {
		r, err = prometheus.GetOneServer(nil,ctx.Params("*"))
		if err!=nil {
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

func AddServer(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	var err error
	server:=new(prometheus.Server)
	DeviceName     ,err := getArg(ctx,"DeviceName"       ,STRING,true ,nil);if err != nil {ctx.JSON(405, err.Error());return}
    DeviceModelId  ,err := getArg(ctx,"DeviceModelId"    ,INT   ,false,nil);if err != nil {ctx.JSON(405, err.Error());return}
    DeviceModelName,err := getArg(ctx,"DeviceModelName"  ,STRING,false,nil);if err != nil {ctx.JSON(405, err.Error());return}
	server.Device.DeviceName=DeviceName.(string)
	server.Device.DeviceModel,err=prometheus.GetOneDeviceModel(DeviceModelId,DeviceModelName)
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	server.Device.GroupId=0
	server.Device.Env=0
	err=prometheus.AddServer(server)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, SUCCESS)
}

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

func UpdateServer(ctx *macaron.Context) {
	var err error=nil
	var server,fake *prometheus.Server
	server, err = prometheus.GetOneServer(nil,ctx.Params("*"))
	if err!=nil{
		ctx.JSON(404, err.Error())
		return
	}
	fake=server.FakeCopy()
	ctx.Req.ParseForm()
    DeviceName,err := getArg(ctx,"DeviceName" ,STRING,false,fake.Get_DeviceName());if err != nil {ctx.JSON(405, err.Error());return}
    GroupId   ,err := getArg(ctx,"GroupId"    ,INT   ,false,fake.Get_GroupId())   ;if err != nil {ctx.JSON(405, err.Error());return}
    Env       ,err := getArg(ctx,"Env"        ,INT   ,false,fake.Get_Env())	      ;if err != nil {ctx.JSON(405, err.Error());return}
    Serial    ,err := getArg(ctx,"Serial" ,STRING,false,fake.Serial)              ;if err != nil {ctx.JSON(405, err.Error());return}
	Hostname  ,err := getArg(ctx,"Hostname" ,STRING,false,fake.Hostname)          ;if err != nil {ctx.JSON(405, err.Error());return}
	Memsize   ,err := getArg(ctx,"Memsize" ,INT,false,fake.Memsize)               ;if err != nil {ctx.JSON(405, err.Error());return}
	Os        ,err := getArg(ctx,"Os" ,STRING,false,fake.Os)                      ;if err != nil {ctx.JSON(405, err.Error());return}
	Release   ,err := getArg(ctx,"Release" ,FLOAT,false,fake.Release)             ;if err != nil {ctx.JSON(405, err.Error());return}

	fake.Device.DeviceName=DeviceName.(string)
	fake.Device.GroupId   =GroupId.(int)
	fake.Device.Env       =uint8(Env.(int))
	fake.Serial    =Serial.(string)
	fake.Hostname  =Hostname.(string)
	fake.Memsize   =uint32(Memsize.(int))
	fake.Os        =Os.(string)
	fake.Release   =Release.(float64)
	err=server.Update(fake)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, err)
}
