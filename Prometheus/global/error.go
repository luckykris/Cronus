package global

type Error string
func (e Error)Error()string{return string(e)}

const (
	ERROR_data_logic            =Error("data logic problem")
	ERROR_device_type_dismatch  =Error("wrong device model for device,because the device model`s device type")
	ERROR_resource_notexist     =Error("resource isn`t exist")
	ERROR_resource_duplicate    =Error("the resource has been declare by other")
	ERROR_resource_outof_range  =Error("resource_outof_range")
	ERROR_parameter_miss        =Error("miss parameter")
)
