package prometheus

import (
	"fmt"
	"testing"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	log "github.com/Sirupsen/logrus"
)

func TestGetAllDeviceType(t *testing.T){
	mainCfg := cfg.LoadCfg()
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Init(mainCfg)
	r,err:=prometheus.GetServer()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n",r)
}
