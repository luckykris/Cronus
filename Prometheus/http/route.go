package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		
		WEB.Group("/server", func() {
			WEB.Get("/?:id:int", GetServer)
			WEB.Post("/", AddServer)
		})
	})
	WEB.NotFound(NotFound)
}
