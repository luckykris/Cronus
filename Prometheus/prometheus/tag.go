package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
	"strings"
)




func DeleteTag(tags ...Tag) error {
	conditions:=[]string{}
	tmp_ls:=[]string{}
	for _,tag:=range tags{
		tmp_ls=append(tmp_ls,string(tag))
	}
	if len(tags)<1{
		return fmt.Errorf("you need specify at least one parameter")
	}
	conditions=append(conditions,fmt.Sprintf(`tag IN ('%s')` ,strings.Join(tmp_ls,`','`)))
	return PROMETHEUS.dbobj.Delete(global.TABLEtag,conditions)
}	


func (self *Device) GetTag() ([]Tag, error) {
	conditions := []string{}
	conditions = append(conditions, fmt.Sprintf("device_id=%d", self.DeviceId))
	var tag string
	r:=[]Tag{}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEdeviceTag, nil,[]string{`tag`}, conditions, &tag)
	if err != nil {
		return r, err
	}
	for cur.Fetch() {
		r = append(r,  Tag(tag))
	}
	return r,nil
}


func (self *Device)AddTag(tags ...Tag) error {
	vals:=[][]interface{}{}
	for _,tag:=range tags{
		vals=append(vals,[]interface{}{self.DeviceId,string(tag)})
	}
	if len(tags)<1{
		return fmt.Errorf("you need specify at least one parameter") 
	}
	return PROMETHEUS.dbobj.Add(global.TABLEdeviceTag, []string{`device_id`, `tag`}, vals)
}


func (self *Device)DeleteTag(tags ...Tag) error {
	conditions:=[]string{fmt.Sprintf("device_id = %d",self.DeviceId)}
	tmp_ls:=[]string{}
	for _,tag:=range tags{
		tmp_ls=append(tmp_ls,string(tag))
	}
	if len(tags)<1{
		return fmt.Errorf("you need specify at least one parameter")
	}
	conditions=append(conditions,fmt.Sprintf(`tag IN ('%s')` ,strings.Join(tmp_ls,`','`)))
	return PROMETHEUS.dbobj.Delete(global.TABLEdeviceTag, conditions)
}
