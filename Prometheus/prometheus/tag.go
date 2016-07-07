package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
	"strings"
)

func GetTag(args ...string) (interface{}, error) {
	var tag string
	var tag_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEtag, []string{`tag_id`, `tag_name`}, args, &tag_id, &tag)
	r := []Tag{}
	for cur.Fetch() {
		r = append(r, Tag{tag_id, tag})
	}
	return r, err
}

func (device *Device) GetTag(id ...int) (interface{}, error) {
	conditions := []string{}
	conditions = append(conditions, fmt.Sprintf("device_id=%d", device.DeviceId))
	if len(id) > 0 {
		conditions = append(conditions, fmt.Sprintf("tag_id=%d", id[0]))
	}
	var tag_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceTag, []string{`tag_id`}, conditions, &tag_id)
	in_tag_ids := []string{}
	for cur.Fetch() {
		in_tag_ids = append(in_tag_ids, fmt.Sprintf("%d", tag_id))
	}
	if err != nil {
		return nil, err
	}
	if len(in_tag_ids) > 0 {
		return GetTag(fmt.Sprintf("tag_id in (%s)", strings.Join(in_tag_ids, ",")))
	} else {
		return []Tag{}, nil
	}
}


func (device *Device)AddTag(tag *Tag) error {
	return PROMETHEUS.dbobj.Add(global.TABLEdeviceTag, []string{`device_id`, `tag_id`}, [][]interface{}{[]interface{}{device.DeviceId,tag.TagId}})
}

func (device *Device)DeleteTag(tag *Tag) error {
	c1 := fmt.Sprintf("tag_id = %d", tag.TagId )
	c2 := fmt.Sprintf("device_id = %d", device.DeviceId)
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceTag, []string{c1,c2})
}