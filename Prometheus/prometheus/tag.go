package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
	"strings"
)

func GetTag(id ...int) ([]Tag, error) {
	var tag string
	var tag_id int
	conditions:=[]string{}
	if len(id)>0{
		tmp_condition:=[]string{}
		for _,v :=range id{
			tmp_condition=append(tmp_condition,fmt.Sprintf("%d",v))
		}
		conditions=append(conditions,fmt.Sprintf("tag_id in (%s)"  ,strings.Join(tmp_condition,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEtag,nil, []string{`tag_id`, `tag_name`}, conditions, &tag_id, &tag)
	r := []Tag{}
	for cur.Fetch() {
		r = append(r, Tag{tag_id, tag})
	}
	return r, err
}
func AddTag(tag *Tag) error {
	return PROMETHEUS.dbobj.Add(global.TABLEtag, []string{`tag_name`}, [][]interface{}{[]interface{}{tag.TagName}})
}
func (tag *Tag) UpdateTag() error {
	c := fmt.Sprintf("tag_id = %d", tag.TagId)
	return PROMETHEUS.dbobj.Update(global.TABLEtag,[]string{c}, []string{`tag_name`}, []interface{}{tag.TagName})
}	
func (tag *Tag) DeleteTag() error {
	c := fmt.Sprintf("tag_id = %d", tag.TagId)
	return PROMETHEUS.dbobj.Delete(global.TABLEtag,[]string{c})
}	


func (device *Device) GetTag(id ...int) (interface{}, error) {
	conditions := []string{}
	conditions = append(conditions, fmt.Sprintf("device_id=%d", device.DeviceId))
	if len(id) > 0 {
		conditions = append(conditions, fmt.Sprintf("tag_id=%d", id[0]))
	}
	var tag_id int
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceTag, nil,[]string{`tag_id`}, conditions, &tag_id)
	in_tag_ids := []int{}
	for cur.Fetch() {
		in_tag_ids = append(in_tag_ids,  tag_id)
	}
	if err != nil {
		return nil, err
	}
	if len(in_tag_ids) > 0 {
		return GetTag(in_tag_ids...)
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
