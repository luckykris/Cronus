package http


const (
	APIVERSION string = "v1"
	server_path string = "/"+APIVERSION+"/"+"server"
)
func LoadRoute() {
	WEB.Group("/"+APIVERSION, func() {
		WEB.Group("/server", func() {
			WEB.Get("/?*", GetServer)
			WEB.Delete("/*", DeleteDevice)
			WEB.Patch("/*", UpdateServer)
			WEB.Post("/", AddServer)
			WEB.Get("/*/space", GetServerSpace)
			WEB.Get("/*/netPorts", GetNetPort)
			WEB.Post("/*/netPorts", AddNetPort)
			WEB.Delete("/*/netPorts", DeleteNetPort)
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
