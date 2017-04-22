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
	"github.com/luckykris/Cronus/Hephaestus/go/net"
)


func GetOneVm(device_id_i interface{},device_name_i interface{})(*Vm,error){
	if device_id_i !=nil{
		s,ok:=DEVICE_INDEX_ID[VM][device_id_i.(int)]
		if ok {
			return s.Value.(*Vm),nil
		}
		return nil,global.ERROR_resource_notexist
	}else if device_name_i!=nil{
		s,ok:=DEVICE_INDEX_NAME[VM][device_name_i.(string)]
		if ok {
			return s.Value.(*Vm),nil
		}
		return nil,global.ERROR_resource_notexist
	}
	return nil,global.ERROR_parameter_miss
}
func GetVm(device_id_i interface{},device_name_i interface{})([]*Vm,error){
	if ifcache(){
		r:=[]*Vm{}
		if device_id_i !=nil{
			s,ok:=DEVICE_INDEX_ID[VM][device_id_i.(int)]
			if ok {
				r=append(r,s.Value.(*Vm))
			}
			return r,nil
		}
		for _,v:=range DEVICE_INDEX_ID[VM]{
			r=append(r,v.Value.(*Vm))
		}
		return r,nil
	}else{
		return GetVmViaDB([]int{} ,[]string{},[]int{} ,[]uint8{})
	}
}

func AddVm(Vm *Vm)(error){
	err:=AddVmViaDB(Vm)
	if err==nil{
		FlushDeviceCache(Vm)
	}
	return err
}

func (Vm *Vm)Update(fake_Vm *Vm)error{
	defer Vm.Unlock()
	Vm.Lock()
	_,err:=Vm.UpdateVmViaDB(fake_Vm)
	if err!=nil{
		return err
	}else{
		DropDeviceCache(Vm)
		FlushDeviceCache(fake_Vm)
		return nil
	}
}

func AddVmViaDB(Vm *Vm)(error) {
	//check 
	_,err:=GetOneVm(nil,Vm.Device.DeviceName)
	if err==nil{
		return global.ERROR_resource_duplicate
	}
	if Vm.Get_DeviceModel().Get_DeviceType() != VM{
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
	rows:=[][]interface{}{[]interface{}{Vm.Get_DeviceName(),
										Vm.Get_DeviceModel().DeviceModelId,
										time.Now().Unix(),
										Vm.Get_GroupId(),
										Vm.Get_Env()}}
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
	conditions:=[]string{fmt.Sprintf("device_name='%s'",Vm.Get_DeviceName())}
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
	Vm.Device.DeviceId=device_id//rewrite Vm struct`s autoincrease id
	err=tx.Add(global.TABLEvm,[]string{`device_id`},[][]interface{}{[]interface{}{device_id}})
	if err!=nil{
		return err
	}
	return err
}


func GetVmViaDB(device_ids []int ,device_names []string,group_ids []int ,envs []uint8)(result []*Vm,err error){
	err = nil
	result=[]*Vm{}
	conditions:=[]string{}
	//device attribute
	var device_id 		  int
	var device_name		  string
	var device_model_id	  int
	var fahter_device_id  int
	var ctime             uint64
	var group_id          int
	var env				  uint8
	//Vm attribute
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
	var mask   uint8
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
							[4]string{db.INNER,global.TABLEvm,strings.Join([]string{global.TABLEvm,"device_id"},"."),strings.Join([]string{global.TABLEnetPort,"device_id"},".")},
					}
	items:=[]string{ `mac`,
					 `ipv4_int`,
					 strings.Join([]string{global.TABLEvm,"device_id"},"."),
					 `mask`,
					 `netPort_type`,
					 `function_type`,
					 `ctime`}
	cur, err := PROMETHEUS.dbobj.GetJoin(global.TABLEnetPort,join_tables_t, nil,items, conditions,  
					&mac,
					&ipv4_int,
					&device_id, 
					&mask,
					&netPort_type,
					&function_type,
					&ctime)
	if err!=nil{
		return 
	}
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
		device_id_map_netports[device_id]=append(device_id_map_netports[device_id], NetPort{ mac_i, net.Ipv4Uint32ConverString(ipv4_int), netPort_type,mask})
	}
	//search Vm
	join_tables_t=[][4]string{
							[4]string{db.LEFT,global.TABLEdevice,strings.Join([]string{global.TABLEdevice,"device_id"},"."),strings.Join([]string{global.TABLEvm,"device_id"},".")},
					}
	items=[]string{strings.Join([]string{global.TABLEvm,"device_id"},"."),
					`device_name`,
					`device_model_id`,
					`fahter_device_id`,
					`ctime`,
					`group_id`,
					`env`,
					`hostname`,
					`memsize`,
					`processor`,
					`os`,
					`_release`,
					`last_change_time`,
					`checksum`}
	cur, err = PROMETHEUS.dbobj.GetJoin(global.TABLEvm,join_tables_t ,nil,items, conditions,  
					&device_id,
					&device_name,
					&device_model_id,
					&fahter_device_id,
					&ctime,
					&group_id,
					&env,
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
		r := new(Vm)
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
		r.Device.FatherDeviceId=fahter_device_id
		r.Device.Ctime=ctime
		r.Device.GroupId=group_id
		r.Device.Env=env
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

func (Vm *Vm)UpdateVmViaDB(fake_Vm *Vm)(*Vm,error) {
	var err error=nil
	conditions:=[]string{fmt.Sprintf("device_id = %d" , Vm.Get_DeviceId())}
	fake_Vm.Checksum=fake_Vm.ComputeSum()
	if fake_Vm.Checksum==Vm.Checksum{
		return Vm,nil
	}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{
		return Vm,err
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
					fake_Vm.Get_DeviceName(),
					fake_Vm.Get_GroupId(),
					fake_Vm.Get_Env(),
				   }
	err=tx.Update(global.TABLEdevice, conditions,items,value)
	if err!=nil{
		return Vm,err
	}
	items2:=[]string{
					 `hostname`, 
					 `memsize`, 
					 `os`,
					 `_release`,
					 `last_change_time`,
					 `checksum`,
					}
	value2:=[]interface{}{
					 	  fake_Vm.Hostname,
					 	  fake_Vm.Memsize,
					 	  fake_Vm.Os,
					 	  fake_Vm.Release,
					 	  time.Now().Unix(),
					 	  fake_Vm.Checksum,
						 }
	err=tx.Update(global.TABLEvm, conditions,items2,value2)
	if err!=nil{
		return Vm,err
	}
	return fake_Vm,err
}

func(Vm *Vm)ComputeSum()string{
	s:=fmt.Sprintf("%s%s%d%s%f%d%d",Vm.Get_DeviceName(),Vm.Hostname,Vm.Memsize,Vm.Os,Vm.Release,Vm.Get_GroupId(),Vm.Get_Env())
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range Vm.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}
func (Vm2 *Vm)FakeCopy()*Vm{
		r := new(Vm)
		r.Device.DeviceId=Vm2.Get_DeviceId()
		r.Device.DeviceName=Vm2.Get_DeviceName()
		r.Device.DeviceModel=Vm2.Get_DeviceModel()
		r.Device.FatherDeviceId=Vm2.Get_FatherDeviceId()
		r.Device.Ctime=Vm2.Device.Ctime
		r.Device.GroupId=Vm2.Get_GroupId()
		r.Device.Env=Vm2.Device.Env
		r.Hostname=Vm2.Hostname
		r.Memsize=Vm2.Memsize
		r.Processor=Vm2.Processor
		r.Os=Vm2.Os
		r.Release=Vm2.Release
		r.LastChangeTime=Vm2.LastChangeTime
		r.Checksum=Vm2.Checksum
		return r
}

//func (Vm *Vm)attr_avaliable()bool{
//	
//}