package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetCabinet(args ...string) (interface{}, error) {
	var id int
	var name string
	var iscloud string
	var capacity_total uint64
	var capacity_used uint64
	var location_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEcabinet, []string{`cabinet_id`, `cabinet_name`, `iscloud`, `capacity_total`, `capacity_used`, `location_id`}, args, &id, &name, &iscloud, &capacity_total, &capacity_used, &location_id)
	r := []global.Cabinet{}
	for cur.Fetch() {
		r = append(r, global.Cabinet{id, name, iscloud, capacity_total, capacity_used, location_id})
	}
	return r, err
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
