package http

import (
	"github.com/go-macaron/macaron"
	//"github.com/macaron-contrib/bindata"
)

var WEB *macaron.Macaron = macaron.New()
var REQ uint64 = 0
const (
	SUCCESS string="success"
)

func Start(debug bool) {
	macaron.Env = macaron.PROD
	_StartMartini(debug)
}

func _StartMartini(debug bool) {
	WEB.Use(macaron.Recovery())
	WEB.Use(macaron.Static("public",
		macaron.StaticOptions{
			//		FileSystem: bindata.Static(bindata.Options{
			//			Asset:      public.Asset,
			//			AssetDir:   public.AssetDir,
			//			AssetNames: public.AssetNames,
			//			Prefix:     "",
			//		}),
			SkipLogging: true,
		},
	))
	WEB.Use(macaron.Renderer(macaron.RenderOptions{
						IndentJSON: debug,
		}))
	//WEB.Use(macaron.Renderer(macaron.RenderOptions{
	//	Directory:  "templates",                //!!when package,escape this line. Specify what path to load the templates from.
	//	Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
	//	Extensions: []string{".tmpl", ".html"}, //!!when package,escape this line.  Specify extensions to load for templates.
	//	Charset:    "UTF-8",                    //!!when package,escape this line.  Sets encoding for json and html content-types. Default is "UTF-8".
	//	//TemplateFileSystem: bindata.Templates(bindata.Options{ //make templates files into bindata.
	//	//	Asset:      templates.Asset,
	//	//	AssetDir:   templates.AssetDir,
	//	//	AssetNames: templates.AssetNames,
	//	//	Prefix:     ""}),
	//}))
	WEB.Use(macaron.Logger())
	LoadRoute()
	WEB.Run("0.0.0.0", 81)
}
