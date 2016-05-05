package http

import (
	"fmt"
	"strconv"
	"github.com/Unknwon/macaron"
)

type CUDRep struct{
	Success bool
	Error interface{}
}

type ArgString struct{
	Name string
	Need bool
}
type ArgInt struct{
	Name string
	Need bool
}

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
func arg2StringOrNil(arg string)(interface{},error){
	var r_arg interface{}
	if arg == "" || arg == "null" {
		r_arg = nil
	} else {
		r_arg =arg
	}
	return r_arg,nil
}

func getAllStringArgs(ctx *macaron.Context,args []ArgString)(map[string]interface{},error){
	args_map:=map[string]interface{}{}
	var err error
	ctx.Req.ParseForm()
	for _,arg:=range args{
		v,err:=arg2StringOrNil(ctx.Req.Form.Get(arg.Name))
		if err!=nil{
			return args_map,err
		}
		if arg.Need && v == nil{
			return args_map,fmt.Errorf("%s can not be nil",arg.Name)
		}
		args_map[arg.Name]=v
	}	
	return args_map,err
}


func getAllIntArgs(ctx *macaron.Context,args []ArgInt)(map[string]interface{},error){
	args_map:=map[string]interface{}{}
	var err error
	ctx.Req.ParseForm()
	for _,arg:=range args{
		v,err:=arg2IntOrNil(ctx.Req.Form.Get(arg.Name))
		if err!=nil{
			return args_map,err
		}
		if arg.Need && v == nil{
			return args_map,fmt.Errorf("%s can not be nil",arg.Name)
		}
		args_map[arg.Name]=v
	}	
	return args_map,err
}
