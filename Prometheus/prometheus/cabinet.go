package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
	log "github.com/Sirupsen/logrus"
)

func GetOneCabinet(id interface{})(*Cabinet, error) {
	var r *Cabinet
	if id!=nil{
		s,ok:=CABINET_INDEX_ID[id.(int)]
		if ok {
			r=s.Value.(*Cabinet)
			return r,nil
		}else{
			return r,global.ERROR_resource_notexist
		}
	}
	return r,nil
}
func GetCabinet(id interface{}) ([]*Cabinet, error) {
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
	var capacity_total uint64
	var capacity_used uint64
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
		cabinet.Idc,tmp_e2=GetOneIdc(idc_id)
		if tmp_e2!=nil{
			log.Error("can`t find idc id:",idc_id)
		}
		r=append(r,cabinet)
	}
	return r,err
}


func AddCabinet(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLEcabinet, []string{`cabinet_name`, `iscloud`, `capacity_total`, `capacity_used`, `location_id`}, values)
}

func DeleteCabinet(id int) error {
	c := fmt.Sprintf("cabinet_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLEcabinet, []string{c})
}

func UpdateCabinet(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("cabinet_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLEcabinet, []string{c}, cloumns, values)
}
