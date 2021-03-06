package net

import (
	"net"
	"fmt"
)

const (
	IPV4MASK uint32 = 1 << 8
	IPV4MIN  uint32 = 1 << 24
)

func Ipv4Uint32ConverString(ipv4_uint32 uint32) string {
	var d_c_b_a = [4]byte{}
	for s := 0; s < 4; s++ {
		d_c_b_a[s] = byte(ipv4_uint32 % IPV4MASK)
		ipv4_uint32 = ipv4_uint32 >> 8
	}
	ipv4 := net.IPv4(d_c_b_a[3], d_c_b_a[2], d_c_b_a[1], d_c_b_a[0])
	return ipv4.String()
}

func Ipv4StringConverUint32(ipv4_str string) (uint32, error) {
	ipv4 := net.ParseIP(ipv4_str)
	if ipv4 == nil {
		return 0, fmt.Errorf("Can't parse ip string :%s ", ipv4_str)
	}
	ipv4 = ipv4.To4()
	var ipv4_uint64 uint32 = 0
	for p := range ipv4 {
		ipv4_uint64 = ipv4_uint64<<8 + uint32(ipv4[p])
	}
	return ipv4_uint64, nil
}

func Ipv4UintAvaliable(ipv4_uint uint32,mask uint)bool{
	if mask  >31 || mask<8 || ipv4_uint <IPV4MIN{
		return false
	}
	if_b_n:=ipv4_uint % (ipv4_uint>>(32-mask))
	if  if_b_n==0 ||if_b_n==1<<(32-mask)-1{
		return false
	}
	return true
}
