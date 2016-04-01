package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"github.com/luckykris/Cronus/Prometheus/global"
	"github.com/luckykris/Cronus/Prometheus/prometheus"
	"strconv"
)

//Device Model
func GetDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetDeviceModel()
	} else {
		r, err = prometheus.GetDeviceModel("device_model_id=" + id)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	switch len(r.([]global.DeviceModel)) {
	case 0:
		ctx.JSON(404, "Not Found")
	case 1:
		r = r.([]global.DeviceModel)[0]
		ctx.JSON(200, &r)
	default:
		ctx.JSON(200, &r)
	}
}

func AddDeviceModel(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("DeviceModelName")
	_type := ctx.Req.Form.Get("DeviceType")
	err := prometheus.AddDeviceModel([][]interface{}{[]interface{}{name, _type}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteDeviceModel(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteDeviceModel(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateDeviceModel(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("DeviceModelName")
	_type := ctx.Req.Form.Get("DeviceType")
	err = _CheckHasClounms("device_model_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("device_type", _type, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateDeviceModel(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

//cabinet
func GetCabinet(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetCabinet()
	} else {
		r, err = prometheus.GetCabinet("cabinet_id=" + id)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	switch len(r.([]global.Cabinet)) {
	case 0:
		ctx.JSON(404, "Not Found")
	case 1:
		r = r.([]global.Cabinet)[0]
		ctx.JSON(200, &r)
	default:
		ctx.JSON(200, &r)
	}
}
func AddCabinet(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("CabinetName")
	iscloud := ctx.Req.Form.Get("IsCloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_total_int, err := strconv.Atoi(capacity_total)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	capacity_used_int, err := strconv.Atoi(capacity_used)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	location_id := ctx.Req.Form.Get("LocationId")
	location_id_int, err := strconv.Atoi(location_id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.AddCabinet([][]interface{}{[]interface{}{name, iscloud, capacity_total_int, capacity_used_int, location_id_int}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteCabinet(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteCabinet(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateCabinet(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("CabinetName")
	iscloud := ctx.Req.Form.Get("IsCloud")
	capacity_total := ctx.Req.Form.Get("CapacityTotal")
	capacity_used := ctx.Req.Form.Get("CapacityUsed")
	location_id := ctx.Req.Form.Get("LocationId")
	err = _CheckHasClounms("cabinet_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("iscloud", iscloud, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("capacity_total", capacity_total, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("capacity_used", capacity_used, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("location_id", location_id, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateCabinet(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

//location
func GetLocation(ctx *macaron.Context) {
	id := ctx.Params("id")
	var r interface{}
	var err error
	if id == "" {
		r, err = prometheus.GetLocation()
	} else {
		r, err = prometheus.GetLocation("location_id=" + id)
	}
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	switch len(r.([]global.Location)) {
	case 0:
		ctx.JSON(404, "Not Found")
	case 1:
		r = r.([]global.Location)[0]
		ctx.JSON(200, &r)
	default:
		ctx.JSON(200, &r)
	}
}
func AddLocation(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("LocationName")
	pic := ctx.Req.Form.Get("Picture")
	father_id := ctx.Req.Form.Get("FatherLocationId")
	var father_id_int interface{}
	var err error
	if father_id == "" || father_id == "null" {
		father_id_int = nil
	} else {
		father_id_int, err = strconv.Atoi(father_id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
	}
	err = prometheus.AddLocation([][]interface{}{[]interface{}{name, pic, father_id_int}})
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(201, "Add Success")
	}
}
func DeleteLocation(ctx *macaron.Context) {
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	err = prometheus.DeleteLocation(id_int)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func UpdateLocation(ctx *macaron.Context) {
	cloumns := []string{}
	values := []interface{}{}
	id := ctx.Params("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(404, "Not Found")
		return
	}
	ctx.Req.ParseForm()
	name := ctx.Req.Form.Get("LocationName")
	pic := ctx.Req.Form.Get("Picture")
	father_id := ctx.Req.Form.Get("FatherLocationId")
	err = _CheckHasClounms("location_name", name, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("picture", pic, false, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = _CheckHasClounms("father_location_id", father_id, true, &cloumns, &values)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	err = prometheus.UpdateLocation(id_int, cloumns, values)
	if err != nil {
		ctx.JSON(400, err.Error())
	} else {
		ctx.JSON(204, "Delete Success")
	}
}

func _CheckHasClounms(cn string, v interface{}, isInt bool, cls *[]string, vals *[]interface{}) error {
	if v == "null" {
		*cls = append(*cls, cn)
		*vals = append(*vals, nil)
	} else if v != "" {
		if isInt {
			v_int, err := strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
			*cls = append(*cls, cn)
			*vals = append(*vals, v_int)
		} else {
			*cls = append(*cls, cn)
			*vals = append(*vals, v)
		}
	}
	return nil
}

func NotFound(ctx *macaron.Context) {
	ctx.JSON(404, "No Such Url")
}
