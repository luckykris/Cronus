package prometheus

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetLocation(name interface{},id ...int)( []*Location, error) {
	locations:=[]*Location{}
	if len(id)!=0 {
		for _,v:=range id{
			locations=append(locations,PROMETHEUS.LocationMapId[v])
		}
		return locations,nil
	}else{
		for _,v:=range PROMETHEUS.LocationMapId{
			locations=append(locations,v)
		}
		return locations,nil
	}
}

func CacheLocation(name interface{},id ...int) error {
	conditions:=[]string{}
	var location_id int
	var location_name string
	var picture string
	var father_id sql.NullInt64
	var father_id_i interface{}
	if name!=nil{
		conditions=append(conditions,fmt.Sprintf("location_name='%s'",name.(string) ) )
	}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("location_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLElocation,nil, []string{`location_id`, `location_name`, `picture`, `father_location_id`}, conditions, &location_id, &location_name, &picture, &father_id)
	if err!=nil{
		return err
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
		location.Picture=picture
		location.FatherLocationId=father_id_i
		PROMETHEUS.LocationMapId[location.LocationId]=location
	}
	return  nil
}


func AddLocation(values [][]interface{}) error {
	return PROMETHEUS.dbobj.Add(global.TABLElocation, []string{`location_name`, `picture`, `father_location_id`}, values)
}

func DeleteLocation(id int) error {
	c := fmt.Sprintf("location_id = %d", id)
	return PROMETHEUS.dbobj.Delete(global.TABLElocation, []string{c})
}

func UpdateLocation(id int, cloumns []string, values []interface{}) error {
	c := fmt.Sprintf("location_id = %d", id)
	return PROMETHEUS.dbobj.Update(global.TABLElocation, []string{c}, cloumns, values)
}
