package os
import (
    "os"
    "os/signal"
    "syscall"
    "runtime"
)



// mask : 0001=once ,0010=async
func StartSignalHandle(signaltype string ,callback func(),mask int){
	if mask&2 ==2 {
		go SignalHandle(signaltype,callback,mask)
		return
	}else{
		SignalHandle(signaltype,callback,mask)
		return
	}
}

func SignalHandle(signaltype string ,callback func(),mask int){
	if mask&2 ==2{
		runtime.Gosched()
	}
	once:=false
	if mask&1 == 1{
		once=true
	}
	for {
		ch := make(chan os.Signal)
 		signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2,syscall.SIGHUP)
 		sig := <-ch
 		v:=sig.String()
 		if v==signaltype{
 				callback()
				if once{
					return
				}
 			}
		}
}
