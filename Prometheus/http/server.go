package http

import (
	"github.com/go-macaron/macaron"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
)


func GetServerSpace(ctx *macaron.Context) {
	var r interface{}
	var err error
	var server *prometheus.Server
	servers, err := prometheus.GetServer(nil,ctx.ParamsInt("id"))
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
	if ctx.Params("id") == ""{
		r, err = prometheus.GetServer(nil,nil)
	} else {
		r, err = prometheus.GetServer(nil,ctx.ParamsInt("id"))
		if len(r.([]*prometheus.Server))<1 {
			ctx.JSON(404, "Not Found")
			return
		}else{
			r = r.([]*prometheus.Server)[0]
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
	var r interface{}
	var err error
	server:=new(prometheus.Server)
	server.Device.DeviceName=ctx.Query("DeviceName")
	server.Device.DeviceModel,err=prometheus.GetDeviceModel(ctx.QueryInt("DeviceModelId"))
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
	ctx.JSON(200, &r)
}

func DeleteServer(ctx *macaron.Context) {
	var err error
	var server *prometheus.Server
	servers, err := prometheus.GetServer(nil,ctx.ParamsInt("id"))
	if len(servers)<1 {
		ctx.JSON(404, "Not Found")
		return
	}else{
		server = servers[0]
	}
	err = server.Delete()
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, nil)
}

func UpdateServer(ctx *macaron.Context) {
	var err error=nil
	var server,fake *prometheus.Server
	server, err = prometheus.GetOneServer(nil,ctx.ParamsInt("id"))
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
	fake.Device.Env       =Env.(uint8)
	fake.Serial    =Serial.(string)
	fake.Hostname  =Hostname.(string)
	fake.Memsize   =Memsize.(uint32)
	fake.Os        =Os.(string)
	fake.Release   =Release.(float64)
	err=server.Update(fake)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, err)
}
