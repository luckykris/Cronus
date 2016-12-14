package prometheus

import (
	//"database/sql"
	//"time"
	//"crypto/md5"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	//log "github.com/Sirupsen/logrus"
)

func GetServer(device_name_i interface{},device_id_i interface{})([]*Server,error){
	r:=[]*Server{}
	if device_id_i !=nil{
		s,ok:=DEVICE_INDEX_ID["server"][device_id_i.(int)]
		if ok {
			r=append(r,s.Value.(*Server))
		}
		return r,nil
	}
	for _,v:=range DEVICE_INDEX_ID["server"]{
		r=append(r,v.Value.(*Server))
	}
	return r,nil
}

func GetServerFromDB(device_ids []int ,device_names []string,group_ids []int ,envs []uint8)(result []*Server,err error){
	err = nil
	result=[]*Server{}
	conditions:=[]string{}
	//device attribute
	var device_id 		  int
	var device_name		  string
	var device_model_id	  int
	var ctime             uint64
	var group_id          int
	var env				  uint8
	//server attribute
	var serial 			  string
	var hostname 		  string
	var memsize 		  uint32
	var processor 		  uint8
	var os				  string
	var release     	  float64
	var last_change_time  uint64
	var checksum          string
	if len(device_ids) >0 {
		conditions=append(conditions,fmt.Sprintf("device_id IN (%s)",int_join(device_ids,",")))
	}
	if len(device_names)>0 {
		conditions=append(conditions,fmt.Sprintf("device_name IN ('%s')",strings.Join(device_names,"','")))
	}
	if len(group_ids)>0 {
		conditions=append(conditions,fmt.Sprintf("group_ids IN (%s)",int_join(group_ids,",")))
	}
	if len(envs)>0 {
		conditions=append(conditions,fmt.Sprintf("env IN (%s)",uint8_join(envs,",")))
	}
	items:=[]string{strings.Join([]string{global.TABLEserver,"device_id"},"."),
					`device_name`,
					`device_model_id`,
					`ctime`,
					`group_id`,
					`env`,
					`serial`,
					`hostname`,
					`memsize`,
					`processor`,
					`os`,
					`_release`,
					`last_change_time`,
					`checksum`}
	cur, err := PROMETHEUS.dbobj.GetLeftJoin(global.TABLEserver,[][3]string{[3]string{global.TABLEdevice,strings.Join([]string{global.TABLEdevice,"device_id"},"."),strings.Join([]string{global.TABLEserver,"device_id"},".")}} ,nil,items, conditions,  
					&device_id,
					&device_name,
					&device_model_id,
					&ctime,
					&group_id,
					&env,
					&serial,
					&hostname,
					&memsize,
					&processor,
					&os,
					&release,
					&last_change_time,
					&checksum)
	for cur.Fetch() {
		r := new(Server)
		r.Device.DeviceId=device_id
		r.Device.DeviceName=device_name
		r.Device.DeviceModel=DEVICEMODEL_INDEX_ID[device_model_id].Value.(*DeviceModel)
		r.Device.FatherDeviceId=nil
		r.Device.Ctime=ctime
		r.Device.GroupId=group_id
		r.Device.Env=env
		r.Serial=serial
		r.Hostname=hostname
		r.Memsize=memsize
		r.Processor=processor
		r.Os=os
		r.Release=release
		r.LastChangeTime=last_change_time
		r.Checksum=checksum
		result=append(result,r)
	}
	return 
}