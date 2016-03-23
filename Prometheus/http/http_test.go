package http

import (
	"fmt"
	"testing"
	"github.com/luckykris/Cronus/Prometheus/http"
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
