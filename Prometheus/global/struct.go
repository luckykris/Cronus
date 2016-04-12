package global

type Device struct {
	DeviceId       int
	DeviceName     string
	DeviceModelId  int
	FatherDeviceId interface{}
	NetPorts       []NetPort
	Tags  []string
}

type NetPort struct {
	Mac      interface{}
	Ipv4     interface{}
	DeviceId int
	Type     string
}

type Space struct {
	CabinetId int
	DeviceId int 
	UPosition int 
	Position string
}

type DeviceModel struct {
	DeviceModelId   int
	DeviceModelName string
	DeviceType      string
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

type Tag struct {
	TagId       int
	Tag     string
}


type DeviceTag struct {
	DeviceId       int
	TagId    int
}