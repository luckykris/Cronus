package prometheus

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/global"
)

//func GetSpace(args ...string) (interface{}, error) {
//	var cabinet_id int
//	var device_id int
//	var u_position int
//	var position string
//	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, args, &cabinet_id, &device_id, &u_position, &position)
//	r := []Space{}
//	for cur.Fetch() {
//		r = append(r, Space{cabinet_id, device_id, u_position, position})
//	}
//	return r, err
//}
//
//
const (
	REAR = "rear"
	FRONT = "front"
)
func (device *Device)GetSpace() ([]Space,error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	conditions:=[]string{fmt.Sprintf("device_id=%d",device.Get_DeviceId())}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, conditions, &cabinet_id, &device_id, &u_position, &position)
	r := []Space{}
	for cur.Fetch() {
		r = append(r, Space{cabinet_id, device_id, u_position, position})
	}
	return r,err
}
func (cabinet *Cabinet)GetSpace() ([]Space,error) {
	var cabinet_id int
	var device_id int
	var u_position int
	var position string
	conditions:=[]string{fmt.Sprintf("cabinet_id=%d",cabinet.CabinetId)}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,nil, []string{`cabinet_id`, `device_id`, `u_position`, `position`}, conditions, &cabinet_id, &device_id, &u_position, &position)
	//init map
	m:= map[string]map[int]Space{}
	for _,v:=range []string{"front","rear"}{
		m[v]=map[int]Space{}
	}
	//
	r := []Space{}
	for cur.Fetch() {
		m[position][u_position]=Space{cabinet_id, device_id, u_position, position}
	}
	for i:=0;i<int(cabinet.CapacityTotal);i++{
		for _,v:=range []string{"front","rear"}{
			s,ok:=m[v][i]
			if ok{
				r=append(r,s)
			}else{
				r=append(r,Space{cabinet.CabinetId, nil, i, v})
			}
		}
	}
	return r,err
}
func GetDeviceCabinetMapViaDB(device_ids []int)(map[int]int,error){
	mm:=map[int]int{}
	conditions:=[]string{}
	var cabinet_id int
	var device_id int
	if len(device_ids)>0{
		conditions=append(conditions,fmt.Sprintf("device_id IN (%s)",int_join(device_ids,",")))
	}
	cur, err := PROMETHEUS.dbobj.Get(global.TABLEspace,"cabinet_id,device_id", []string{`cabinet_id`, `device_id`}, conditions, &cabinet_id, &device_id)
	if err!=nil{
		return mm,err
	}
	for cur.Fetch() {
		mm[device_id]=cabinet_id
	}
	return mm, err
}

func (device *Device)SetSpace(cabinet_id ,start_u uint32,front_rear string)error{
	cabinet,err:=GetOneCabinet(cabinet_id,nil)
	if err!=nil{
		return err
	}
	deviceModel:=device.Get_DeviceModel()
	//check if start_u is avaliable
	var end_u=start_u+uint32(deviceModel.U)-1
	if start_u > cabinet.CapacityTotal || end_u >cabinet.CapacityTotal {
		return global.ERROR_resource_outof_range
	}
	tx,err:=PROMETHEUS.dbobj.Begin()
	if err!=nil{
		return err
	}
	defer func(){
		if err!=nil{
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	//check if the space is free
	conditions:=[]string{fmt.Sprintf("cabinet_id=%d",cabinet_id),
					 	 fmt.Sprintf("u_position>=%d",start_u),
					 	 fmt.Sprintf("u_position<=%d",end_u),
				}
	if deviceModel.HALF_FULL == HALF {
		conditions=append(conditions,fmt.Sprintf("position=%s"),front_rear)
	}
	var inuse_space int 
	cur,err := tx.Get(global.TABLEspace,nil,[]string{"count(1)"},conditions,&inuse_space)
	if !cur.Fetch(){
		err = global.ERROR_resource_notexist
	}
	cur.Close()
	if err!=nil{
		return err
	}
	if inuse_space >0{
		return global.ERROR_resource_duplicate
	}
	take_spaces:=[][]interface{}{}
	switch deviceModel.HALF_FULL{ 
	case HALF:
		for i:=start_u;i<=end_u;i++{
			take_spaces=append(take_spaces,[]interface{}{cabinet_id,device.Get_DeviceId(),i,front_rear})
		}
	case FULL:
		for i:=start_u;i<=end_u;i++{
			take_spaces=append(take_spaces,[]interface{}{cabinet_id,device.Get_DeviceId(),i,FRONT})
			take_spaces=append(take_spaces,[]interface{}{cabinet_id,device.Get_DeviceId(),i,REAR})
		}
	}
	err=tx.Add(global.TABLEspace,[]string{`cabinet_id`,`device_id`,`u_position`,`position`},take_spaces)
	return err
}

func (device *Device)DeleteSpace()(error) {
	conditions:=[]string{fmt.Sprintf("device_id=%d",device.DeviceId)}
	return PROMETHEUS.dbobj.Delete(global.TABLEspace, conditions)
}

//
//
