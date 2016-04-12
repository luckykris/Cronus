package prometheus

import (
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetTag(args ...string) (interface{}, error) {
	var tag string
	var tag_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEtag, []string{`tag_id`, `tag`}, args, &tag_id, &tag)
	r := []global.Tag{}
	for cur.Fetch() {
		r = append(r, global.Tag{tag_id,tag})
	}
	return r, err
}
