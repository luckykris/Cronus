package prometheus

import (
	"fmt"
)

func int_join(i_ls []int,spliter string)string{
	tmp_s:=""
	for i:= range i_ls{
		tmp_s+=fmt.Sprintf("%s%s",i_ls[i],spliter)
	}
	if len(i_ls)<1{
		return tmp_s
	}else{
		return tmp_s[0:len(i_ls)-1]
	}
}
func uint8_join(i_ls []uint8,spliter string)string{
	tmp_s:=""
	for i:= range i_ls{
		tmp_s+=fmt.Sprintf("%s%s",i_ls[i],spliter)
	}
	if len(i_ls)<1{
		return tmp_s
	}else{
		return tmp_s[0:len(i_ls)-1]
	}
}