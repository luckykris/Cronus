package prometheus

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetOneLocation(id interface{},name interface{})(*Location, error) {
	var r *Location
	if id!=nil{
		s,ok:=LOCATION_INDEX_ID[id.(int)]
		if ok {
			r=s.Value.(*Location)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}else if name!=nil{
		s,ok:=LOCATION_INDEX_NAME[name.(string)]
		if ok {
			r=s.Value.(*Location)
			return r,nil
		}
		return r,global.ERROR_resource_notexist
	}
	return nil,global.ERROR_parameter_miss
}
func GetLocation(id interface{},name interface{})( []*Location, error) {
	r:=[]*Location{}
	if id!=nil{
		s,ok:=LOCATION_INDEX_ID[id.(int)]
		if ok {
			r=append(r,s.Value.(*Location))
			return r,nil
		}else{
			return r,global.ERROR_resource_notexist
		}
	}
	for _,v:=range LOCATION_INDEX_ID{
		r=append(r,v.Value.(*Location))
	}
	return r,nil
}

func GetLocationViaDB(location_ids []int,location_names []string)([]*Location,error){
	r:=[]*Location{}
	conditions:=[]string{}
	var location_id int
	var err error
	var location_name string
	var father_id sql.NullInt64
	var father_id_i interface{}
	if len(location_names)>0{
		conditions=append(conditions,fmt.Sprintf(`location_name IN ('%s')`,strings.Join(location_names,"','")))
	}
	if len(location_ids)>0{
		conditions=append( conditions,fmt.Sprintf(`location_id IN (%s)`,int_join(location_ids,",")))
	}
	items:=[]string{`location_id`, 
				    `location_name`, 
				    `father_location_id`}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLElocation,nil, items, conditions, &location_id,
																				  &location_name,
																				  &father_id)
	if err!=nil{
		return r,err
	}
	for cur.Fetch() {
		if !father_id.Valid {
			father_id_i = nil
		} else {
			father_id_i = int(father_id.Int64)
		}
		location:=new(Location)
		location.LocationId=location_id
		location.LocationName=location_name
		location.FatherLocationId=father_id_i
		r=append(r,location)
	}
	return  r,err
}

func AddLocation(self *Location)error{
	err:=AddLocationViaDB(self)
	if err==nil{
		create_cache_and_index(self)
	}
	return err
}

func AddLocationViaDB(self *Location) error {
	_,err:=GetOneLocation(nil,self.LocationName)
	if err==nil{return global.ERROR_resource_duplicate}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{return err}
	items:=[]string{`location_name`, 
			        `father_location_id`}
	values:=[][]interface{}{[]interface{}{
					self.LocationName,
					self.FatherLocationId,
	}}
	err=tx.Add(global.TABLElocation, items, values)
	if err!=nil{return err}
	defer func(){if err!=nil{tx.Rollback()}else{tx.Commit()}}()
	var id int
	conditions:=[]string{fmt.Sprintf("location_name='%s'",self.LocationName)}
	items2:=[]string{"location_id"}
	cur,err:=tx.Get(global.TABLElocation, nil,items2, conditions,  
					&id)
	if !cur.Fetch(){
		err= global.ERROR_data_logic
	}
	cur.Close()
	self.LocationId=id
	return err
}

func (self *Location)Delete()(err error){
	defer self.Unlock()
	self.Lock()
	err=self.DeleteViaDB()
	if err==nil{
		drop_cache_and_index(self)
	}
	return 
}
func (self *Location)DeleteViaDB()error{
	conditions:=[]string{}
	conditions=append(conditions,fmt.Sprintf("location_id=%d",self.LocationId))
	return PROMETHEUS.dbobj.Delete(global.TABLElocation,conditions)
}
func (self *Location)Update(fake *Location)(err error){
	defer self.Unlock()
	self.Lock()
	drop_cache_and_index(self)
	err=self.UpdateViaDB(fake)
	create_cache_and_index(self)
	return 
}

func (self *Location)UpdateViaDB(fake *Location)(error) {
	c := fmt.Sprintf("location_id = %d", self.LocationId)
	items:=[]string{`location_name`,
					}
	values:=[]interface{}{
					fake.LocationName,
						}
	err:=PROMETHEUS.dbobj.Update(global.TABLElocation, []string{c}, items, values)
	if err!=nil{return err}
	self.LocationName =fake.LocationName
	return err
}

func (self *Location)FakeCopy()*Location{
		r := new(Location)
		r.LocationId=self.LocationId
		r.LocationName =self.LocationName 
		r.FatherLocationId=self.FatherLocationId
		return r
}