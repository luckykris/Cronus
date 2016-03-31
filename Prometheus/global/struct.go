package global

type DeviceModel struct {
	Id   int
	Name string
	Type string
}

type Cabinet struct {
	Id   int
	Name string
	IsCloud string
	CapacityTotal int64
	CapacityUsed int64
	LocationId   int
}

type Location struct{
	Id int
	Name string
	Pic string
	FatherId int
}
