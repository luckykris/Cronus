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
			WEB.Get("/?*", GetLocation)
			WEB.Delete("/*", DeleteLocation)
			WEB.Patch("/*", UpdateLocation)
			WEB.Post("/", AddLocation)
		})
		WEB.Group("/idc", func() {
			WEB.Get("/?*", GetIdc)
			WEB.Delete("/*", DeleteIdc)
			WEB.Patch("/*", UpdateIdc)
			WEB.Post("/", AddIdc)
		})
		WEB.Group("/cabinet", func() {
			WEB.Get("/?*", GetCabinet)
			WEB.Delete("/*", DeleteCabinet)
			WEB.Patch("/*", UpdateCabinet)
			WEB.Post("/", AddCabinet)
			WEB.Get("/*/space", GetCabinetSpace)
		})
		WEB.Group("/deviceModel", func() {
			WEB.Get("/?*", GetDeviceModel)
			WEB.Post("/", AddDeviceModel)
			WEB.Delete("/*", DeleteDeviceModel)
		})
	})
	WEB.NotFound(NotFound)
}
