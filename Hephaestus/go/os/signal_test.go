package os

import (
	"fmt"
	"github.com/luckykris/Cronus/Hephaestus/os"
	"testing"
)

func TestMain(t *testing.T) {
	os.StartSignalHandle("interrupt",cb,0)
	fmt.Printf("i finish mask . \n")
	//select{}
}

func cb(){
	fmt.Printf("i got signal. \n")
}
