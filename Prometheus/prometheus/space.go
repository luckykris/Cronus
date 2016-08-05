package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

func GetSpace(args ...string) (interface{}, error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, args, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r, err
}


func (device *Device)GetSpace() ([]Space,error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	conditions:=[]string{fmt.Sprintf("device_id=%d",device.DeviceId)}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, conditions, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r,err
}


func (device *Device)AddSpace(cabinet_id ,start_u int,position string)(error){
	fmt.Printf("%#v \n",device.DeviceModel)
	var err error
	spaces:=[][]interface{}{}
	end_u:=start_u+device.DeviceModel.U
	//check if the space is used
	r,err:=GetSpace(fmt.Sprintf("cabinet_id=%d",cabinet_id),fmt.Sprintf("position='%s'",position),fmt.Sprintf("u_position >=%d",start_u),fmt.Sprintf("u_position <=%d",end_u))
	if err!=nil{
		return err
	}
	if len(r.([]Space)) >0{
		return fmt.Errorf("space is used by other device!")
	}
	if end_u > int(PROMETHEUS.CabinetMapId[cabinet_id].CapacityTotal) {
		return fmt.Errorf("out of cabinet's Capacity ")
	}
	err=device.DeleteSpace()
	if err!=nil{
		return fmt.Errorf("reset space failed :%s",err.Error())
	}
	for i:= start_u;i<end_u;i++{
		spaces=append(spaces,[]interface{}{cabinet_id,device.DeviceId,i,"front"})
	}
	err=PROMETHEUS.dbobj.Add(global.TABLEspace, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, spaces)
	return err
}

func (device *Device)DeleteSpace()(error) {
	conditions:=[]string{fmt.Sprintf("device_id=%d",device.DeviceId)}
	return PROMETHEUS.dbobj.Delete(global.TABLEspace, conditions)
}



func (cabinet *Cabinet)GetSpace() ([]Space,error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	conditions:=[]string{fmt.Sprintf("cabinet_id=%d",cabinet.CabinetId)}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, conditions, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r,err
}