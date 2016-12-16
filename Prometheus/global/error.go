package global

type Error string
func (e Error)Error()string{return string(e)}

const (
	ERROR_data_logic         =Error("data logic problem")
	ERROR_resource_notexist  =Error("resource isn`t exist")
	ERROR_resource_duplicate =Error("the resource has been declare by other")
)
