package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		
		WEB.Group("/server", func() {
			WEB.Get("/?:id:int", GetServer)
			WEB.Delete("/?:id:int", DeleteServer)
			WEB.Patch("/?:id:int", UpdateServer)
			WEB.Post("/", AddServer)
		})
	})
	WEB.NotFound(NotFound)
}
