package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		WEB.Group("/server", func() {
			WEB.Get("/?:id:int", GetServer)
			WEB.Get("/:id:int/space", GetServerSpace)
			WEB.Delete("/?:id:int", DeleteServer)
			WEB.Patch("/?:id:int", UpdateServer)
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
			WEB.Get("/?:id:int", GetDeviceModel)
		})
	})
	WEB.NotFound(NotFound)
}
