package net

import (
	"fmt"
	"github.com/luckykris/Cronus/Hephaestus/go/net"
	"testing"
)

func TestIpv4StringConverUint32(t *testing.T) {
	fmt.Println(net.Ipv4StringConverUint32("192.168.33.71"))
	fmt.Println(net.Ipv4Uint32ConverString(3232244039))
	fmt.Println(net.Ipv4Uint32ConverString(3232244039))
	ip1,_:=net.Ipv4StringConverUint32("192.168.33.0")
	ip2,_:=net.Ipv4StringConverUint32("192.168.33.255")
	ip3,_:=net.Ipv4StringConverUint32("192.168.33.254")
	ip4,_:=net.Ipv4StringConverUint32("0.168.33.254")
	ip5,_:=net.Ipv4StringConverUint32("256.168.33.254")
	fmt.Println(net.Ipv4UintAvaliable(ip1,24))//f 
	fmt.Println(ip2)//f
	fmt.Println(net.Ipv4UintAvaliable(ip2,24))//f
	fmt.Println(net.Ipv4UintAvaliable(ip4,24))//f
	fmt.Println(net.Ipv4UintAvaliable(ip5,24))//f
	fmt.Println(net.Ipv4UintAvaliable(ip3,24))//t
}
