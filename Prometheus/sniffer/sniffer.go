package sniffer
import (
	"strings"
)

type Sniffer uint8

const (
	None	Sniffer = iota
	Ansible 
)

func ParseSniffer(sniffer string) (Sniffer, error) {
	switch strings.ToLower(sniffer) {
	case "none":
		return None,nil
	case "ansible":
		return Ansible, nil
	}

	var s Sniffer
	return s, fmt.Errorf("unknow sniffer: %q", sniffer)
}