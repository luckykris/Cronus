package http

func LoadRoute() {
	WEB.Group("/v1", func() {
		WEB.Group("/deviceType", func(){
			 WEB.Get("/?:id:int", GetDevice)
		})
	})
//	WEB.NotFound(NotFound)
}
