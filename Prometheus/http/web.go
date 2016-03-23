package http

import (
	"github.com/Unknwon/macaron"
	//"github.com/macaron-contrib/bindata"
	"runtime"
	"time"
)

var WEB *macaron.Macaron = macaron.New()
var REQ uint64 = 0

func Start() {
	_StartMartini()
}

func _StartMartini() {
	WEB.Use(macaron.Recovery())
	WEB.Use(macaron.Static("public",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      public.Asset,
				AssetDir:   public.AssetDir,
				AssetNames: public.AssetNames,
				Prefix:     "",
			}),
			SkipLogging: true,
		},
	))

	WEB.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "templates",                //!!when package,escape this line. Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, //!!when package,escape this line.  Specify extensions to load for templates.
		Charset:    "UTF-8",                    //!!when package,escape this line.  Sets encoding for json and html content-types. Default is "UTF-8".
		//TemplateFileSystem: bindata.Templates(bindata.Options{ //make templates files into bindata.
		//	Asset:      templates.Asset,
		//	AssetDir:   templates.AssetDir,
		//	AssetNames: templates.AssetNames,
		//	Prefix:     ""}),
	}))
	LoadRoute()
	WEB.Run("0.0.0.0", 80)
}
