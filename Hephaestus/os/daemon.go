package os
import (
        "os"
        "path/filepath"
)
func Daemonize(){
        if os.Getppid()!=1{
                filePath,_:=filepath.Abs(os.Args[0])
                args:=append([]string{filePath},os.Args[1:]...)
                os.StartProcess(filePath,args,&os.ProcAttr{Files:[]*os.File{os.Stdin,os.Stdout,os.Stderr}})
		os.Exit(0)
        }
}
