package http

import (
	"fmt"
	"strconv"
	"github.com/go-macaron/macaron"
)


const (
		INT =iota
		STRING 
		FLOAT
)
func arg2IntOrNil(arg string)(interface{},error){
	var r_arg interface{}
	var err error = nil
	if arg == "" || arg == "null" {
		r_arg = nil
	} else {
		r_arg, err= strconv.Atoi(arg)
	}
	return r_arg,err
}
func arg2StringOrNil(arg string)(interface{}){
	var r_arg interface{}
	if arg == "" || arg == "null" {
		r_arg = nil
	} else {
		r_arg =arg
	}
	return r_arg
}

func getArg(ctx *macaron.Context,key string ,t int,need bool,defaultv interface{})(interface{},error){
	var err error =nil
	var r interface{}
	switch t{
	case STRING:
		r=arg2StringOrNil(ctx.Req.Form.Get(key))
	case INT:
		r,err=arg2IntOrNil(ctx.Req.Form.Get(key))
		if err!=nil{
			return r,fmt.Errorf("'%s' must be int.",key)
		}
	}
	if r==nil{
		if need{
			return r,fmt.Errorf("'%s' must be specify and must not be null.",key)
		}else{
			return defaultv,err
		} 
	}else{
		return r,err
	}
}


