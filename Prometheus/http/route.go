package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		WEB.Group("/deviceModel", func() {
			WEB.Get("/?:id:int", GetDeviceModel)
			WEB.Post("/", AddDeviceModel)
			WEB.Delete("/:id:int", DeleteDeviceModel)
			WEB.Patch("/:id:int", UpdateDeviceModel)
		})
		WEB.Group("/cabinet", func() {
			WEB.Get("/?:id:int", GetCabinet)
			WEB.Post("/", AddCabinet)
			WEB.Delete("/:id:int", DeleteCabinet)
			WEB.Patch("/:id:int", UpdateCabinet)
			WEB.Get("/:CabinetId:int/space", GetCabinetSpace)
		})
		WEB.Group("/location", func() {
			WEB.Get("/?:id:int", GetLocation)
			WEB.Post("/", AddLocation)
			WEB.Delete("/:id:int", DeleteLocation)
			WEB.Patch("/:id:int", UpdateLocation)
		})
		WEB.Group("/idc", func() {
			WEB.Get("/?:id:int", GetIdc)
		})
		WEB.Group("/tag", func() {
			WEB.Delete("/:tag:string", DeleteTag)
		})
		WEB.Group("/device", func() {
			WEB.Get("/?:id:int", GetDevice)
			WEB.Post("/", AddDevice)
			WEB.Delete("/:id:int", DeleteDevice)
			//WEB.Patch("/:id:int", UpdateDevice)
			WEB.Get("/:DeviceId:int/netPort/?:NetPortId:int", GetNetPort)
			WEB.Post("/:DeviceId:int/netPort/", AddNetPort)
			WEB.Patch("/:DeviceId:int/netPort/?:NetPortId:int", UpdateNetPort)
			WEB.Delete("/:DeviceId:int/netPort/?:NetPortId:int", DeleteNetPort)
			WEB.Get("/:DeviceId:int/tag/", GetDeviceTag)
			WEB.Post("/:DeviceId:int/tag/:Tag:string", AddDeviceTag)
			WEB.Delete("/:DeviceId:int/tag/:Tag:string", DeleteDeviceTag)
			WEB.Get("/:DeviceId:int/space", GetDeviceSpace)
			WEB.Post("/:DeviceId:int/space", AddDeviceSpace)
			WEB.Delete("/:DeviceId:int/space", DeleteDeviceSpace)
		})
		WEB.Group("/server", func() {
			WEB.Get("/?:id:int", GetServer)
			WEB.Post("/", AddServer)
		})
		WEB.Group("/vm", func() {
			WEB.Get("/?:id:int", GetVm)
			WEB.Post("/", AddVm)
		})
		WEB.Group("/space", func() {
			WEB.Get("/", GetSpace)
		})
	})
	WEB.NotFound(NotFound)
}
