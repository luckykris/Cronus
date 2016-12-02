package net

import (
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/toolkit"
	"testing"
)

func TestIpv4StringConverUint32(t *testing.T) {
	fmt.Println(toolkit.Ipv4StringConverUint32("192.168.33.71"))
	fmt.Println(toolkit.Ipv4Uint32ConverString(3232244039))
}
