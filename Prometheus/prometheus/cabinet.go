package prometheus

import (
	"fmt"
	"strings"
	"time"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)

func GetOneCabinet(id interface{},name interface{})(*Cabinet, error) {
	var r *Cabinet
	if id!=nil{
		s,ok:=CABINET_INDEX_ID[id.(int)]
		if ok {
			r=s.Value.(*Cabinet)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}else if name!=nil{
		s,ok:=CABINET_INDEX_NAME[name.(string)]
		if ok {
			r=s.Value.(*Cabinet)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}
	return nil,global.ERROR_parameter_miss
}
func GetCabinet(id interface{},name interface{})( []*Cabinet, error) {
	r:=[]*Cabinet{}
	if id!=nil{
		s,ok:=CABINET_INDEX_ID[id.(int)]
		if ok {
			r=append(r,s.Value.(*Cabinet))
			return r,nil
		}else{
			return r,global.ERROR_resource_notexist
		}
	}
	for _,v:=range CABINET_INDEX_ID{
		r=append(r,v.Value.(*Cabinet))
	}
	return r,nil
}

func GetCabinetViaDB(ids []int,names []string)([]*Cabinet,error){
	r:=[]*Cabinet{}
	conditions:=[]string{}
	var cabinet_id int
	var cabinet_name string
	var capacity_total uint32
	var capacity_used uint32
	var idc_id int
	if len(names)>0{
		conditions=append(conditions,fmt.Sprintf(`cabinet_name IN ('%s')`,strings.Join(names,"','")))
	}
	if len(ids)>0{
		conditions=append( conditions,fmt.Sprintf(`cabinet_id IN (%s)`,int_join(ids,",")))
	}
	items:=[]string{`cabinet_id`, 
				    `cabinet_name`, 
				    `capacity_total`, 
				    `capacity_used`, 
				    `idc_id`}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEcabinet,nil, items, conditions, &cabinet_id, 
																				 &cabinet_name, 
																				 &capacity_total, 
																				 &capacity_used, 
																				 &idc_id)
	if err!=nil{
		return r,err
	}
	var tmp_e2 error
	for cur.Fetch() {
		cabinet:=new(Cabinet)
		cabinet.CabinetId=cabinet_id
		cabinet.CabinetName=cabinet_name
		cabinet.CapacityTotal=capacity_total
		cabinet.CapacityUsed=capacity_used
		cabinet.Idc,tmp_e2=GetOneIdc(idc_id,nil)
		if tmp_e2!=nil{
			log.Error("can`t find idc id:",idc_id)
		}
		r=append(r,cabinet)
	}
	return r,err
}


func AddCabinet(self *Cabinet)error{
	err:=AddCabinetViaDB(self)
	if err==nil{
		create_cache_and_index(self)
	}
	return err
}
func AddCabinetViaDB(self *Cabinet) error {
	_,err:=GetOneCabinet(nil,self.CabinetName)
	if err==nil{return global.ERROR_resource_duplicate}
	tx,err:=PROMETHEUS.dbobj.Begin()
	items:=[]string{`cabinet_name`,
					`capacity_total`,
					`capacity_used`,
					`idc_id`,
					`ctime`,
					}
	values:=[][]interface{}{[]interface{}{
					self.CabinetName,
					self.CapacityTotal,
					0,//self.CapacityUsed,
					self.Idc.IdcId,
					time.Now().Unix(),
		}}
	err=tx.Add(global.TABLEcabinet, items, values)
	if err!=nil{return err}
	defer func(){if err!=nil{tx.Rollback()}else{tx.Commit()}}()
	var id int
	conditions:=[]string{fmt.Sprintf("cabinet_name='%s'",self.CabinetName)}
	items2:=[]string{"cabinet_id"}
	cur,err:=tx.Get(global.TABLEcabinet, nil,items2, conditions,  
					&id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	cur.Close()
	self.CabinetId=id
	return err
}

func (self *Cabinet)Delete()(err error){
	defer self.Unlock()
	self.Lock()
	err=self.DeleteViaDB()
	if err==nil{
		drop_cache_and_index(self)
	}
	return 
}
func (self *Cabinet)DeleteViaDB()error{
	conditions:=[]string{}
	conditions=append(conditions,fmt.Sprintf("cabinet_id=%d",self.CabinetId))
	return PROMETHEUS.dbobj.Delete(global.TABLEcabinet,conditions)
}

func (self *Cabinet)Update(fake *Cabinet)(err error){
	defer self.Unlock()
	self.Lock()
	drop_cache_and_index(self)
	err=self.UpdateViaDB(fake)
	create_cache_and_index(self)
	return 
}
func (self *Cabinet)UpdateViaDB(fake *Cabinet)(error) {
	c := fmt.Sprintf("cabinet_id = %d", self.CabinetId)
	items:=[]string{`cabinet_name` ,
					`capacity_used`,}
	values:=[]interface{}{
					fake.CabinetName,
					fake.CapacityUsed,
		}
	err:=PROMETHEUS.dbobj.Update(global.TABLEcabinet, []string{c}, items, values)
	if err!=nil{return err}
	self.CabinetName =fake.CabinetName
	self.CapacityUsed=fake.CapacityUsed
	return err
}
func (self *Cabinet)FakeCopy()*Cabinet{
		r := new(Cabinet)
		r.CabinetId=self.CabinetId
		r.CabinetName =self.CabinetName 
		r.CapacityTotal=self.CapacityTotal
		r.CapacityUsed=self.CapacityUsed
		r.Idc=self.Idc
		return r
}