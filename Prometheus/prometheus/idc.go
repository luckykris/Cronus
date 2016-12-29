package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)

func GetOneIdc(id interface{},name interface{})(*Idc, error) {
	var r *Idc
	if id!=nil{
		s,ok:=IDC_INDEX_ID[id.(int)]
		if ok {
			r=s.Value.(*Idc)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}else if name!=nil{
		s,ok:=IDC_INDEX_NAME[name.(string)]
		if ok {
			r=s.Value.(*Idc)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}
	return nil,global.ERROR_parameter_miss
}
func GetIdc(id interface{},name interface{})( []*Idc, error) {
	r:=[]*Idc{}
	if id!=nil{
		s,ok:=IDC_INDEX_ID[id.(int)]
		if ok {
			r=append(r,s.Value.(*Idc))
			return r,nil
		}else{
			return r,global.ERROR_resource_notexist
		}
	}
	for _,v:=range IDC_INDEX_ID{
		r=append(r,v.Value.(*Idc))
	}
	return r,nil
}

func GetIdcViaDB(idc_ids []int,idc_names []string)([]*Idc ,error) {
	r:=[]*Idc{}
	conditions:=[]string{}
	var idc_id int
	var idc_name string
	var location_id int
	var err error=nil
	if len(idc_names)>0{
		conditions=append(conditions,fmt.Sprintf(`idc_name IN ('%s')`,strings.Join(idc_names,"','")))
	}
	if len(idc_ids)>0{
		conditions=append( conditions,fmt.Sprintf(`idc_id IN (%s)`,int_join(idc_ids,",")))
	}
	items:=[]string{`idc_id`, `idc_name`, `location_id`}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEidc,nil,items, conditions, &idc_id,
																			&idc_name, 
																			&location_id)
	if err!=nil{
		return r,err
	}
	var tmp_e2 error
	for cur.Fetch() {
		idc:=new(Idc)
		idc.IdcId=idc_id
		idc.IdcName=idc_name
		idc.Location,tmp_e2=GetOneLocation(location_id,nil)
		if tmp_e2!=nil{
			log.Error("can`t find location id:",location_id)
		}
		r=append(r,idc)
	}
	return  r,err
}


func AddIdc(self *Idc)error{
	err:=AddIdcViaDB(self)
	if err==nil{
		create_cache_and_index(self)
	}
	return err
}

func AddIdcViaDB(self *Idc) error {
	_,err:=GetOneIdc(nil,self.IdcName)
	if err==nil{return global.ERROR_resource_duplicate}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{return err}
	items:=[]string{`idc_name`, 
			        `location_id`}
	values:=[][]interface{}{[]interface{}{
					self.IdcName,
					self.Location.LocationId,
	}}
	err=tx.Add(global.TABLEidc, items, values)
	if err!=nil{return err}
	defer func(){if err!=nil{tx.Rollback()}else{tx.Commit()}}()
	var id int
	conditions:=[]string{fmt.Sprintf("idc_name='%s'",self.IdcName)}
	items2:=[]string{"idc_id"}
	cur,err:=tx.Get(global.TABLEidc, nil,items2, conditions,  
					&id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	cur.Close()
	self.IdcId=id
	return err
}

func (self *Idc)Delete()(err error){
	defer self.Unlock()
	self.Lock()
	err=self.DeleteViaDB()
	if err==nil{
		drop_cache_and_index(self)
	}
	return 
}
func (self *Idc)DeleteViaDB()error{
	conditions:=[]string{}
	conditions=append(conditions,fmt.Sprintf("idc_id=%d",self.IdcId))
	return PROMETHEUS.dbobj.Delete(global.TABLEidc,conditions)
}

func (self *Idc)Update(fake *Idc)(err error){
	defer self.Unlock()
	self.Lock()
	drop_cache_and_index(self)
	err=self.UpdateViaDB(fake)
	create_cache_and_index(self)
	return 
}

func (self *Idc)UpdateViaDB(fake *Idc)(error) {
	c := fmt.Sprintf("idc_id = %d", self.IdcId)
	items:=[]string{`idc_name`,
					}
	values:=[]interface{}{
					fake.IdcName,
						}
	err:=PROMETHEUS.dbobj.Update(global.TABLEidc, []string{c}, items, values)
	if err!=nil{return err}
	self.IdcName =fake.IdcName
	return err
}

func (self *Idc)FakeCopy()*Idc{
		r := new(Idc)
		r.IdcId=self.IdcId
		r.IdcName =self.IdcName 
		r.Location=self.Location
		return r
}