package prometheus

import (
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetCabinet(name interface{},id ...int) (interface{}, error) {
	cabinets:=[]*Cabinet{}
	if len(id) !=0 {
		for _,v:=range id{
			cabinets=append(cabinets,PROMETHEUS.CabinetMapId[v])
		}
		return cabinets,nil
	}else{
		for _,v:=range PROMETHEUS.CabinetMapId{
			cabinets=append(cabinets,v)
		}
		return cabinets,nil
	}
}

func CacheCabinet(name interface{},id ...int)error{
	conditions:=[]string{}
	var cabinet_id int
	var cabinet_name string
	var iscloud string
	var capacity_total uint64
	var capacity_used uint64
	var location_id int
	if name!=nil{
		conditions=append(conditions,fmt.Sprintf("cabinet_name='%s'",name.(string)))
	}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("cabinet_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEcabinet,nil, []string{`cabinet_id`, `cabinet_name`, `iscloud`, `capacity_total`, `capacity_used`, `location_id`}, conditions, &cabinet_id, &cabinet_name, &iscloud, &capacity_total, &capacity_used, &location_id)
	for cur.Fetch() {
		cabinet:=new(Cabinet)
		cabinet.CabinetId=cabinet_id
		cabinet.CabinetName=cabinet_name
		cabinet.IsCloud=iscloud
		cabinet.CapacityTotal=capacity_total
		cabinet.CapacityUsed=capacity_used
		cabinet.LocationId=location_id
		if _,ok:=PROMETHEUS.CabinetMapId[cabinet.CabinetId];ok{
			delete(PROMETHEUS.CabinetMapId,cabinet.CabinetId)
		}
		PROMETHEUS.CabinetMapId[cabinet.CabinetId]=cabinet
	}
	return err
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
