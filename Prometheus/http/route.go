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
		})
		WEB.Group("/location", func() {
			WEB.Get("/?:id:int", GetLocation)
			WEB.Post("/", AddLocation)
			WEB.Delete("/:id:int", DeleteLocation)
			WEB.Patch("/:id:int", UpdateLocation)
		})
		WEB.Group("/tags", func() {
			WEB.Get("/?:id:int", GetTag)
			WEB.Post("/", AddTag)
			WEB.Delete("/:id:int", DeleteTag)
			WEB.Patch("/:id:int", UpdateTag)
		})
		WEB.Group("/devices", func() {
			WEB.Get("/?:id:int", GetDevice)
			WEB.Post("/", AddDevice)
			WEB.Delete("/:id:int", DeleteDevice)
			WEB.Patch("/:id:int", UpdateDevice)
			WEB.Get("/:DeviceId:int/netPorts/?:NetPortId:int", GetNetPort)
			WEB.Post("/:DeviceId:int/netPorts/", AddNetPort)
			WEB.Patch("/:DeviceId:int/netPorts/?:NetPortId:int", UpdateNetPort)
			WEB.Delete("/:DeviceId:int/netPorts/?:NetPortId:int", DeleteNetPort)
			WEB.Get("/:DeviceId:int/tags/?:TagId:int", GetDeviceTag)
			WEB.Post("/:DeviceId:int/tags/:TagId:int", AddDeviceTag)
			WEB.Delete("/:DeviceId:int/tags/:TagId:int", DeleteDeviceTag)
		})
		WEB.Group("/space", func() {
			WEB.Get("/", GetSpace)
		})
	})
	WEB.NotFound(NotFound)
}
