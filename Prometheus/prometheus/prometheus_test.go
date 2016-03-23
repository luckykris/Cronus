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
	r,err:=prometheus.GetAllDeviceType()
	if err!=nil{
		return
	}
	fmt.Println(r)
}
