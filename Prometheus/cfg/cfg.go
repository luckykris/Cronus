package cfg

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"os"
)

const (
	SOFTWARE string = "prometheus"
)

var LOGLEVELMAP = map[string]log.Level{"debug": log.DebugLevel, "info": log.InfoLevel, "warning": log.WarnLevel, "error": log.ErrorLevel, "fatal": log.FatalLevel, "panic": log.PanicLevel}

type MainCfg struct {
	DbCfg  DbCfgStruct  `toml:"database"`
	LogCfg LogCfgStruct `toml:"log"`
	SnifferCfg SnifferCfgStruct `toml:"sniffer"`
	Daemon bool
	Debug  bool
}

//database config struct for toml
type DbCfgStruct struct {
	Class        string `toml:"class"`
	Host         string `toml:"host"`
	Port         int64  `toml:"port"`
	Db           string `toml:"db"`
	Charset      string `toml:"charset"`
	User         string `toml:"user"`
	Passwd       string `toml:"passwd"`
	MaxLifeTime  int64  `toml:"maxLifeTime"`
	MaxIdleConns int64  `toml:"maxIdleConns"`
}

//log config struct for toml
type LogCfgStruct struct {
	Level   string    `toml:"level"`
	LevelId log.Level `toml:"levelId"`
}

//log config struct for toml
type SnifferCfgStruct struct {
	Class   	string  `toml:"class"`
	PluginPath  string 	`toml:"plugin_path"`
	Interval	int64 	`toml:"interval"`
}


//Load all config
func LoadCfg() MainCfg {
	debug:=flag.Bool("debug", false, "Start in debug mode.")
	daemon:=flag.Bool("daemon", false, "Start in daemon mode.")
	cfgfile := flag.String("config", "/etc/"+SOFTWARE+".toml", "Configuration file ")
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
		fmt.Printf("Configuration Error:%s\n", err.Error())
		os.Exit(-1)
	}
	//sys config
	mainCfgObj.Daemon = *daemon
	mainCfgObj.Debug = *debug
	//config  database
	if meta.IsDefined("database") {
		if !meta.IsDefined("database", "maxLifeTime") {
			mainCfgObj.DbCfg.MaxLifeTime = 0
		}
		if !meta.IsDefined("database", "maxIdleConns") {
			mainCfgObj.DbCfg.MaxIdleConns = 5
		}
		if !meta.IsDefined("database", "class") {
			mainCfgObj.DbCfg.Class = "mysql"
		}
		if !meta.IsDefined("database", "port") {
			switch mainCfgObj.DbCfg.Class {
			case "mysql":
				mainCfgObj.DbCfg.Port = 3306
			default:
				fmt.Printf("Do not support DB type %s\n", mainCfgObj.DbCfg.Class)
				os.Exit(-1)
			}
		}
	} else {
		fmt.Printf("DataBase has not been defined.\n")
		os.Exit(-1)
	}
	//

	//config log
	if !meta.IsDefined("log", "level") {
		mainCfgObj.LogCfg.Level = "info"
	}
	levelId, err := log.ParseLevel(mainCfgObj.LogCfg.Level)
	if err != nil {
		fmt.Printf(err.Error()+"\n")
		os.Exit(-1)
	} else {
		mainCfgObj.LogCfg.LevelId = levelId
	}
	//config promethues 
	if !meta.IsDefined("sniffer", "class") {
		mainCfgObj.SnifferCfg.Class = "none"
	}
	if !meta.IsDefined("sniffer", "interval") {
		mainCfgObj.SnifferCfg.Interval = 1
	}
	if !meta.IsDefined("sniffer", "plugin_path") {
		mainCfgObj.SnifferCfg.PluginPath = "/usr/share/" + SOFTWARE
	}
	return mainCfgObj
}
