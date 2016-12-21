package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetIdc(name interface{},id ...int)( []*Idc, error) {
	r:=[]*Idc{}
	if len(id)!=0 {
		for _,v:=range id{
			r=append(r,PROMETHEUS.IdcMapId[v])
		}
		return r,nil
	}else{
		for _,v:=range PROMETHEUS.IdcMapId{
			r=append(r,v)
		}
		return r,nil
	}
}

func CacheIdc(name interface{},id ...int) error {
	conditions:=[]string{}
	var idc_id int
	var idc_name string
	var location_id int
	if name!=nil{
		conditions=append(conditions,fmt.Sprintf("idc_name='%s'",name.(string) ) )
	}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("idc_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEidc,nil, []string{`idc_id`, `idc_name`, `location_id`}, conditions, &idc_id, &idc_name, &location_id)
	if err!=nil{
		return err
	}
	for cur.Fetch() {
		idc:=new(Idc)
		idc.IdcId=idc_id
		idc.IdcName=idc_name
		idc.LocationId=location_id
		PROMETHEUS.IdcMapId[idc.IdcId]=idc
	}
	return  nil
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
