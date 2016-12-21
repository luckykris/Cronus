package prometheus

import (
	"database/sql"
	"fmt"
	"strings"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetLocation(location_id interface{})( []*Location, error) {
	r:=[]*Location{}
	if location_id!=nil{
		s,ok:=LOCATION_INDEX_ID[location_id.(int)]
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
