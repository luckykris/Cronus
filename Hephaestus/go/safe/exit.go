package safe
import "sync"

var SAFE chan struct{} = make(chan struct{})
var ALL sync.WaitGroup

type Exiter struct{
	ExitSignal chan struct{}
	WaitExitSafely sync.WaitGroup
}

func NewExit()*Exiter{
	var s chan struct{} = make(chan struct{})
	return &Exiter{ExitSignal:s}
}

func (self *Exiter)StartAllExit() {
	close(self.ExitSignal)
}
func (self *Exiter)WaitExitSignal() chan struct{} {
	return self.ExitSignal
}
func (self *Exiter)FinishOneExit() {
	self.WaitExitSafely.Done()
}

func (self *Exiter)Join(){
	self.WaitExitSafely.Add(1)
}
func (self *Exiter)WaitAllExit(){
	self.WaitExitSafely.Wait()
}
