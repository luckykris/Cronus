package os
import (
        "os"
        "path/filepath"
)
func Daemonize(){
        if os.Getppid()!=1{
                filePath,_:=filepath.Abs(os.Args[0])
                args:=append([]string{filePath},os.Args[1:]...)
		devNull,_ := os.OpenFile(os.DevNull, os.O_RDWR, 0)
                os.StartProcess(filePath,args,&os.ProcAttr{Files:[]*os.File{devNull,devNull,devNull}})
		os.Exit(0)
        }
}
