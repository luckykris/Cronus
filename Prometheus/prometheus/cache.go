package prometheus
import (
	//"github.com/luckykris/Cronus/Prometheus/global"
	//"database/sql"
)




func LoadCache()error{
	var err error
	err=LoadDeviceModel()
	err=LoadServer()
	err=LoadVm()
	err=LoadCabinet()
	err=LoadLocation()
	err=LoadIdc()
	return err
}


func LoadServer() error {
	err:=CacheServer(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}

func LoadVm() error {
	err:=CacheVm(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}


func LoadDeviceModel()error{
	err:=CacheDeviceModel(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}

func LoadCabinet()error{
	err:=CacheCabinet(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}
func LoadLocation()error{
	err:=CacheLocation(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}
func LoadIdc()error{
	err:=CacheIdc(nil) 
	if err!=nil{
		return err
	}else{
		return nil
	}
}