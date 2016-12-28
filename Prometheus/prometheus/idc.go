package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)

func GetOneIdc(idc_id interface{})(*Idc, error) {
	var r *Idc
	if idc_id!=nil{
		s,ok:=IDC_INDEX_ID[idc_id.(int)]
		if ok {
			r=s.Value.(*Idc)
			return r,nil
		}else{
			return r,global.ERROR_resource_notexist
		}
	}
	return r,nil
}
func GetIdc(idc_id interface{})( []*Idc, error) {
	r:=[]*Idc{}
	if idc_id!=nil{
		s,ok:=IDC_INDEX_ID[idc_id.(int)]
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


func AddIdc(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLEidc, []string{`idc_name`, `location_id`}, values)
}

func DeleteIdc(id int) error {
	c := fmt.Sprintf("idc_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEidc, []string{c})
}

func UpdateIdc(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("idc_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLEidc, []string{c}, cloumns, values)
}
