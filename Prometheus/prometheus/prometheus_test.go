package prometheus

import (
	"fmt"
	"testing"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	log "github.com/Sirupsen/logrus"
)

func TestPrometheus(t *testing.T){
	mainCfg := cfg.LoadCfg()
	log.SetLevel(mainCfg.LogCfg.LevelId)
	prometheus.Init(mainCfg)
	r,err:=prometheus.GetDeviceModelFromDB([]string{},[]string{},[]int{})
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n",r)
}
