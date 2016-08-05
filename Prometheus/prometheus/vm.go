package prometheus

import (
	//"database/sql"
	"time"
	"crypto/md5"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	//log "github.com/Sirupsen/logrus"
)


func GetVm(name interface{},id ...int) ([]*Vm, error){
	vms:=[]*Vm{}
	if len(id) !=0 {
		for _,v:=range id{
			vms=append(vms,PROMETHEUS.VmMapId[v])
		}
		return vms,nil
	}else{
		for _,v:=range PROMETHEUS.VmMapId{
			vms=append(vms,v)
		}
		return vms,nil
	}
}
func CacheVm(name interface{},id ...int) (error) {
	var device_id int
	var hostname string
	var memsize int
	var os string
	var release float64
	var last_change_time int64
	var checksum string
	conditions:=[]string{}
	devices,err:=GetDevice(name,id...)
	device_map:=map[int]*Device{}
	if err!=nil{
		return err
	}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("device_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEvm, nil, []string{`device_id`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, conditions, &device_id, &hostname, &memsize,&os,&release,&last_change_time,&checksum)
	if err != nil {
		return  err
	}
	for _,device:=range devices {
		device_map[device.DeviceId] = device
	}
	for cur.Fetch() {
		if _,ok:=device_map[device_id];!ok{
			return fmt.Errorf("device data fatal!")
		}
		vm:=new(Vm)
		vm.Device.NetPorts=device_map[device_id].NetPorts
		vm.Device.DeviceId=device_id
		vm.Device.DeviceName=device_map[device_id].DeviceName
		vm.Device.DeviceModel=device_map[device_id].DeviceModel
		vm.Device.FatherDeviceId=device_map[device_id].FatherDeviceId
		vm.Hostname=hostname
		vm.Memsize=memsize
		vm.Os=os
		vm.Release=release
		vm.LastChangeTime=last_change_time
		vm.Checksum=checksum
		PROMETHEUS.DeviceMapId[vm.Device.DeviceId]=vm
		PROMETHEUS.VmMapId[vm.Device.DeviceId]=vm
	}
	return err
}



func AddVm(device_name string,device_model_id ,fatherDeviceId int) error {
	err:=AddDevice(device_name,device_model_id,fatherDeviceId)
	if err!=nil{
		return err
	}
	devices,err:=GetDevice(device_name)
	if err!=nil{
		return err
	}
	if len(devices)!=1{
		return fmt.Errorf("amazing critical error!")
	}
	err=PROMETHEUS.dbobj.Add(global.TABLEvm, []string{`device_id`, `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, [][]interface{}{[]interface{}{devices[0].DeviceId,"Unknow",0,"Unknow",0,0,"Never"}})
	if err!=nil{
		return err
	}else{
		vm:=new(Vm)
		vm.Device.DeviceId=devices[0].DeviceId
		vm.Device.DeviceName=devices[0].DeviceName
		vm.Device.DeviceModel=devices[0].DeviceModel
		vm.Device.FatherDeviceId=devices[0].FatherDeviceId
		vm.Hostname="Unknow"
		vm.Memsize=0
		vm.Os="Unknow"
		vm.Release=0
		vm.LastChangeTime=0
		vm.Checksum="Never"
		PROMETHEUS.DeviceMapId[vm.Device.DeviceId]=vm
		PROMETHEUS.VmMapId[vm.Device.DeviceId]=vm
		return nil
	}
}


func (self *Vm)UpdateVm() error {
	var err error
	conditions:=[]string{fmt.Sprintf("device_id = %d" , self.DeviceId)}
	server_pre_ls,err:=GetVm(nil,self.DeviceId)
	if err!=nil{
		return err
	}
	checksum:=self.ComputSum()
	if server_pre_ls[0].Checksum == checksum{
		return nil
	}
	self.LastChangeTime=time.Now().Unix()
	err=PROMETHEUS.dbobj.Update(global.TABLEvm, conditions, []string{ `hostname`, `memsize`, `os`,`release`,`last_change_time`,`checksum`}, []interface{}{self.Hostname,self.Memsize,self.Os,self.Release,self.LastChangeTime,checksum})
	if err!=nil{
		return err
	}
	err=CacheServer(nil,self.Device.DeviceId)
	if err!=nil{
		return err
	}
	return nil
}


func(self *Vm)ComputSum()string{
	s:=fmt.Sprintf("%s%d%s%f",self.Hostname,self.Memsize,self.Os,self.Release)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	//netPorts_count_ipv4_int:=0
	//for _,netPort :=range server.NetPorts{
//		netPorts_count_ipv4_int+=netPort.Ipv4Int
//	}
}

