package diskq

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	a, b := New("diskq_test.test", "", 5, 10)
	if b != nil {
		fmt.Println(b)
	}
	a.Put([]byte(`string`),false)
	a.Put([]byte(`1111`),false)
	a.Put([]byte(`2222`),false)
	//fd.Write(c)
	c, _ := a.Get(false)
	if string(c) == string([]byte(`string`)) {
		t.Logf("ok")
	}
}
