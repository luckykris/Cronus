package safe

import (
	"fmt"
	"github.com/luckykris/Cronus/Hephaestus/safe"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	exit:=safe.NewExit()
	fmt.Println("init exit")
	go childExitProcess(1,exit)
	go childExitProcess(2,exit)
	time.Sleep(2*time.Second)
	fmt.Println("start exiting")
	StartAllExit()
	WaitExitSignal()
	
}

func childExitProcess(child_no int,e safe.Exiter){
	fmt.Printf("i am child %d,i begin to Run. \n",child_no)
	e.Join()
	select{
		case <-e.WaitExitSignal():
			fmt.Printf("i am child %d,i begin to Exit. \n",child_no)
			time.Sleep(2*time.Second)
			e.FinishOneExit()
			fmt.Printf("i am child %d,i finish Exiting. \n ",child_no)
	}
}
