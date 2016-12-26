package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		WEB.Group("/server", func() {
			WEB.Get("/?*", GetServer)
			WEB.Get("/*/space", GetServerSpace)
			WEB.Delete("/*", DeleteServer)
			WEB.Patch("/*", UpdateServer)
			WEB.Post("/", AddServer)
		})
		WEB.Group("/location", func() {
			WEB.Get("/?:id:int", GetLocation)
		})
		WEB.Group("/idc", func() {
			WEB.Get("/?:id:int", GetIdc)
		})
		WEB.Group("/cabinet", func() {
			WEB.Get("/?:id:int", GetCabinet)
			WEB.Get("/:id:int/space", GetCabinetSpace)
		})
		WEB.Group("/deviceModel", func() {
			WEB.Get("/?*", GetDeviceModel)
			WEB.Post("/", AddDeviceModel)
			WEB.Delete("/*", DeleteDeviceModel)
		})
	})
	WEB.NotFound(NotFound)
}
