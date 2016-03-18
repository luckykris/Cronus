package cfg

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var CFGFILE *string //new cfg way

const (
	SOFTWARE string = "prometheus"
)

type MainCfg struct {
	DbCfg DbCfgStruct `toml:"DataBase"`
}

type DbCfgStruct struct {
	Class   string `toml:"class"`
	Host    string `toml:"host"`
	Db      string `toml:"db"`
	Charset string `toml:"charset"`
	User    string `toml:"user"`
	Passwd  string `toml:"passwd"`
}

func LoadCfg() MainCfg {
	cfgfile := flag.String("config", "/etc/"+SOFTWARE+".toml", "Configuration file path")
	help := flag.Bool("help", false, "Show all the help infomation")
	version := flag.Bool("version", false, "Show version")
	flag.Parse()
	if *help {
		fmt.Println("====================================")
		fmt.Println("==============" + SOFTWARE + "==============")
		fmt.Println("====================================")
		fmt.Println("Usage:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *version {
		fmt.Printf("Version:%f%d \n", VERSION, RELEASE)
		os.Exit(0)
	}
	var mainCfgObj MainCfg
	meta, err := toml.DecodeFile(*cfgfile, &mainCfgObj)
	if err != nil {
		fmt.Printf("Configuration Error:%s", err.Error())
		os.Exit(-1)
	}
	if meta.IsDefined(SOFTWARE, "DataBase") {

	}

	return mainCfgObj
}
