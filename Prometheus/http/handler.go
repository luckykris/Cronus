package http

import (
	"github.com/Unknwon/macaron"
	"net/http"
)

func (ctx *macaron.Context) GetDevice(r macaron.Render, req *http.Request) {
	r.JSON(200, []string{ctx.Data["id"]})
}

func NotFound(r macaron.Render) {
	r.HTML(404, "notfound", nil)
}
