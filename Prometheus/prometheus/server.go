package prometheus

import (
	"database/sql"
	"crypto/md5"
	"time"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/db"
	log "github.com/Sirupsen/logrus"
)


func GetServer(device_name_i interface{},device_id_i interface{})([]*Server,error){
	r:=[]*Server{}
	if device_id_i !=nil{
		s,ok:=DEVICE_INDEX_ID[SERVER][device_id_i.(int)]
		if ok {
			r=append(r,s.Value.(*Server))
		}
		return r,nil
	}
	for _,v:=range DEVICE_INDEX_ID[SERVER]{
		r=append(r,v.Value.(*Server))
	}
	return r,nil
}

func AddServer(server *Server)(error){
	err:=AddServerViaDB(server)
	if err!=nil{
		return err
	}else{
		FlushDeviceCache(server)
		return nil
	}
}

func AddServerViaDB(server *Server)(error) {
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{
		return err
	}
	defer func(){
		if err!=nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	//check
	if if_device_name_exist(server.Device.DeviceName){
		return global.ERROR_resource_duplicate
	}
	//
	rows:=[][]interface{}{[]interface{}{server.Device.DeviceName,
										server.Device.DeviceModel.DeviceModelId,
										time.Now().Unix(),
										server.Device.GroupId,
										server.Device.Env}}
	err=tx.Add(global.TABLEdevice,[]string{`device_name`,
									  `device_model_id`,
									  `ctime`,
									  `group_id`,
									  `env`,
									  },rows)
	if err!=nil{
		return err
	}
	var device_id int
	conditions:=[]string{fmt.Sprintf("device_name='%s'",server.Device.DeviceName)}
	items:=[]string{"device_id"}
	cur,err:=tx.Get(global.TABLEdevice, nil,items, conditions,  
					&device_id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	cur.Close()
	if err!=nil{
		return err
	}
	err=tx.Add(global.TABLEserver,[]string{`device_id`},[][]interface{}{[]interface{}{device_id}})
	if err!=nil{
		return err
	}
	return err
}


func GetServerViaDB(device_ids []int ,device_names []string,group_ids []int ,envs []uint8)(result []*Server,err error){
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
	//netPort attribute
	var mac sql.NullString
	var mac_i interface{}
	var ipv4_int uint32
	var netPort_type string
	var function_type string

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
	//get netport
	join_tables_t:=[][4]string{
							[4]string{db.INNER,global.TABLEserver,strings.Join([]string{global.TABLEserver,"device_id"},"."),strings.Join([]string{global.TABLEnetPort,"device_id"},".")},
					}
	items:=[]string{ `mac`,
					 `ipv4_int`,
					 strings.Join([]string{global.TABLEserver,"device_id"},"."),
					 `netPort_type`,
					 `function_type`,
					 `ctime`}
	cur, err := PROMETHEUS.dbobj.GetJoin(global.TABLEnetPort,join_tables_t, nil,items, conditions,  
					&mac,
					&ipv4_int,
					&device_id, 
					&netPort_type,
					&function_type,
					&ctime)
	device_id_map_netports := map[int][]NetPort{}
	for cur.Fetch() {
		if !mac.Valid {
			mac_i = nil
		} else {
			mac_i = mac.String
		}
		_,ok:=device_id_map_netports[device_id]
		if !ok{
			device_id_map_netports[device_id]=[]NetPort{}
		}
		device_id_map_netports[device_id]=append(device_id_map_netports[device_id], NetPort{ mac_i, ipv4_int, netPort_type})
	}
	//search server
	join_tables_t=[][4]string{
							[4]string{db.LEFT,global.TABLEdevice,strings.Join([]string{global.TABLEdevice,"device_id"},"."),strings.Join([]string{global.TABLEserver,"device_id"},".")},
					}
	items=[]string{strings.Join([]string{global.TABLEserver,"device_id"},"."),
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
	cur, err = PROMETHEUS.dbobj.GetJoin(global.TABLEserver,join_tables_t ,nil,items, conditions,  
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
	var tmp_e2 error //tmp variable error
	for cur.Fetch() {
		r := new(Server)
		r.Device.DeviceId=device_id
		r.Device.DeviceName=device_name
		r.Device.DeviceModel,tmp_e2=GetDeviceModel(device_model_id)
		if tmp_e2!=nil{
			log.Debug("can`t find device model id:",device_model_id)
		}
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
		tmp_netPorts,ok:=device_id_map_netports[device_id]
		if ok{
			r.NetPorts = tmp_netPorts
		}else{
			r.NetPorts = []NetPort{}
		}
		result=append(result,r)
	}
	return 
}

func UpdateServerViaDB(server *Server) error {
	var err error
	conditions:=[]string{fmt.Sprintf("device_id = %d" , server.DeviceId)}
	server_pre_ls,err:=GetServer(nil,server.DeviceId)
	if err!=nil{
		return err
	}
	checksum:=server.ComputSum()
	if server_pre_ls[0].Checksum == checksum{
		return nil
	}
	server.LastChangeTime=uint64(time.Now().Unix())
	err=PROMETHEUS.dbobj.Update(global.TABLEserver, conditions, []string{`serial`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, []interface{}{server.Serial,server.Hostname,server.Memsize,server.Os,server.Release,server.LastChangeTime,checksum})
	if err!=nil{
		return err
	}
	if err!=nil{
		return err
	}
	return nil
}
//func (self *Server)Update(d Device)(err error){
//	defer self.Unlock()
//	self.Lock()
//	err=nil
//	err=self.UpdateViaDB(d)
//	if err!=nil{
//		return 
//	}
//	FlushDeviceCache(self)
//	return
//}
//func (self *Server)UpdateViaDB(d Device)(err error){
//	err = nil
//	conditions:=[]string{fmt.Sprintf("device_id = %d" , self.DeviceId)}
//	items:=[]string{`device_name`,
//					`ctime`,
//					`group_id`,
//					`env`,
//					}
//	value:=[]interface{}{d.DeviceName,
//					   time.Now().Unix(),
//					   d.GroupId,
//					   d.Env,
//					  }
//	err=PROMETHEUS.dbobj.Update(global.TABLEdevice, conditions, items,value)
//	return 
//}

func(server *Server)ComputSum()string{
	s:=fmt.Sprintf("%s%s%d%s%f",server.Serial,server.Hostname,server.Memsize,server.Os,server.Release)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range server.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}