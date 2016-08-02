package prometheus
import (
	//"github.com/luckykris/Cronus/Prometheus/global"
	//"database/sql"
)




func LoadCache()error{
	var err error
	err=LoadDeviceModel()
	err=LoadServer()
	err=LoadCabinet()
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