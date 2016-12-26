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


func GetOneServer(device_id_i interface{},device_name_i interface{})(*Server,error){
	if device_id_i !=nil{
		s,ok:=DEVICE_INDEX_ID[SERVER][device_id_i.(int)]
		if ok {
			return s.Value.(*Server),nil
		}
		return nil,global.ERROR_resource_notexist
	}else if device_name_i!=nil{
		s,ok:=DEVICE_INDEX_NAME[SERVER][device_name_i.(string)]
		if ok {
			return s.Value.(*Server),nil
		}
		return nil,global.ERROR_resource_notexist
	}
	return nil,global.ERROR_parameter_miss
}
func GetServer(device_id_i interface{},device_name_i interface{})([]*Server,error){
	if ifcache(){
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
	}else{
		return GetServerViaDB([]int{} ,[]string{},[]int{} ,[]uint8{})
	}
}

func AddServer(server *Server)(error){
	err:=AddServerViaDB(server)
	if err==nil{
		FlushDeviceCache(server)
	}
	return err
}

func (server *Server)Update(fake_server *Server)error{
	defer server.Unlock()
	server.Lock()
	_,err:=server.UpdateServerViaDB(fake_server)
	if err!=nil{
		return err
	}else{
		DropDeviceCache(server)
		FlushDeviceCache(fake_server)
		return nil
	}
}

func AddServerViaDB(server *Server)(error) {
	//check 
	if if_device_name_exist(server.Device.DeviceName){
		return global.ERROR_resource_duplicate
	}
	if server.Get_DeviceModel().Get_DeviceType() != SERVER{
		return global.ERROR_device_type_dismatch
	}
	//

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
	//
	rows:=[][]interface{}{[]interface{}{server.Get_DeviceName(),
										server.Get_DeviceModel().DeviceModelId,
										time.Now().Unix(),
										server.Get_GroupId(),
										server.Get_Env()}}
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
	conditions:=[]string{fmt.Sprintf("device_name='%s'",server.Get_DeviceName())}
	items:=[]string{"device_id"}
	cur,err:=tx.Get(global.TABLEdevice, nil,items, conditions,  
					&device_id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	cur.Close()
	server.Device.DeviceId=device_id//rewrite server struct`s autoincrease id
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
	//get cabinet
	device_id_map_cabinet_id,err:=GetDeviceCabinetMapViaDB(device_ids)
	if err!=nil{
		return 
	}
	//
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
	var ok bool
	for cur.Fetch() {
		r := new(Server)
		r.Device.DeviceId=device_id
		r.Device.DeviceName=device_name
		r.Device.DeviceModel,tmp_e2=GetOneDeviceModel(device_model_id,nil)
		if tmp_e2!=nil{
			log.Error("can`t find device model id:",device_model_id)
		}
		r.Device.CabinetId,ok=device_id_map_cabinet_id[device_id]
		if !ok{
			r.Device.CabinetId=nil
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

func (server *Server)UpdateServerViaDB(fake_server *Server)(*Server,error) {
	var err error=nil
	conditions:=[]string{fmt.Sprintf("device_id = %d" , server.Get_DeviceId())}
	fake_server.Checksum=fake_server.ComputeSum()
	if fake_server.Checksum==server.Checksum{
		return server,nil
	}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{
		return server,err
	}
	defer func(){
		if err!=nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	items:=[]string{`device_name`,
					`group_id`,
					`env`,
				   }
	value:=[]interface{}{
					fake_server.Get_DeviceName(),
					fake_server.Get_GroupId(),
					fake_server.Get_Env(),
				   }
	err=tx.Update(global.TABLEdevice, conditions,items,value)
	if err!=nil{
		return server,err
	}
	items2:=[]string{`serial`, 
					 `hostname`, 
					 `memsize`, 
					 `os`,
					 `_release`,
					 `last_change_time`,
					 `checksum`,
					}
	value2:=[]interface{}{fake_server.Serial,
					 	  fake_server.Hostname,
					 	  fake_server.Memsize,
					 	  fake_server.Os,
					 	  fake_server.Release,
					 	  time.Now().Unix(),
					 	  fake_server.Checksum,
						 }
	err=tx.Update(global.TABLEserver, conditions,items2,value2)
	if err!=nil{
		return server,err
	}
	return fake_server,err
}

func(server *Server)ComputeSum()string{
	s:=fmt.Sprintf("%s%s%s%d%s%f%d%d",server.Get_DeviceName(),server.Serial,server.Hostname,server.Memsize,server.Os,server.Release,server.Get_GroupId(),server.Get_Env())
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range server.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}
func (server *Server)FakeCopy()*Server{
		r := new(Server)
		r.Device.DeviceId=server.Get_DeviceId()
		r.Device.DeviceName=server.Get_DeviceName()
		r.Device.DeviceModel=server.Get_DeviceModel()
		r.Device.FatherDeviceId=nil
		r.Device.Ctime=server.Device.Ctime
		r.Device.GroupId=server.Get_GroupId()
		r.Device.Env=server.Device.Env
		r.Serial=server.Serial
		r.Hostname=server.Hostname
		r.Memsize=server.Memsize
		r.Processor=server.Processor
		r.Os=server.Os
		r.Release=server.Release
		r.LastChangeTime=server.LastChangeTime
		r.Checksum=server.Checksum
		return r
}

//func (server *Server)attr_avaliable()bool{
//	
//}