package sniffer
import (
	"strings"
	"fmt"
	"github.com/luckykris/Cronus/Prometheus/cfg"
	"time"
	"path/filepath"
)


const (
	AnsiblePlugin string = "ansible-sniff.py"
)

type Snifferi interface{
	Start()
}

func Init(cfg cfg.SnifferCfgStruct) (Snifferi, error) {
	switch strings.ToLower(cfg.Class) {
	case "none":
		return &None{Interval:time.Duration(cfg.Interval)},nil
	case "ansible":
		return &Ansible{Interval:time.Duration(cfg.Interval),Exe:filepath.Join(cfg.PluginPath,AnsiblePlugin)},nil
	default:
		return nil, fmt.Errorf("unknow sniffer: %q", cfg.Class)
	}
}

