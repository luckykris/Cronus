package global

type Device struct {
	DeviceId int 
	DeviceName string
	DeviceModelId int 
	FatherDeviceId interface{}
}

type DeviceModel struct {
	DeviceModelId   int
	DeviceModelName string
	DeviceType string
}

type Cabinet struct {
	CabinetId     int
	CabinetName   string
	IsCloud       string
	CapacityTotal uint64
	CapacityUsed  uint64
	LocationId    int
}

type Location struct {
	LocationId       int
	LocationName     string
	Picture          string
	FatherLocationId interface{}
}
