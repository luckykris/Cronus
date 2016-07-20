package prometheus

import (
	"database/sql"
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetLocation(args ...string) (interface{}, error) {
	var id int
	var name string
	var picture string
	var father_id sql.NullInt64
	cur, err := PROMETHEUS.dbobj.Get(global.TABLElocation,nil, []string{`location_id`, `location_name`, `picture`, `father_location_id`}, args, &id, &name, &picture, &father_id)
	r := []Location{}
	for cur.Fetch() {
		if !father_id.Valid {
			r = append(r, Location{id, name, picture, nil})
		} else {
			r = append(r, Location{id, name, picture, int(father_id.Int64)})
		}
	}
	return r, err
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
