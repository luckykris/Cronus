package http

import (
	//"fmt"
	"github.com/Unknwon/macaron"
	"strconv"
)







func _CheckHasClounms(cn string, v interface{}, isInt bool, cls *[]string, vals *[]interface{}) error {
	if v == "null" {
		*cls = append(*cls, cn)
		*vals = append(*vals, nil)
	} else if v != "" {
		if isInt {
			v_int, err := strconv.Atoi(v.(string))
			if err != nil {
				return err
			}
			*cls = append(*cls, cn)
			*vals = append(*vals, v_int)
		} else {
			*cls = append(*cls, cn)
			*vals = append(*vals, v)
		}
	}
	return nil
}

func NotFound(ctx *macaron.Context) {
	ctx.JSON(404, "No Such Url")
}
