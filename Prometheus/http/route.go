package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		WEB.Group("/deviceType", func(){
			 m.Get("/:id", GetDevice)
		}
	})
	WEB.NotFound(NotFound)
}
