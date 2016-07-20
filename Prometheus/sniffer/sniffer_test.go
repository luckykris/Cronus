package sniffer

import (
        "fmt"
        "testing"
        "github.com/luckykris/Cronus/Prometheus/sniffer"
)

func TestAnsible(t *testing.T){
        //an:=&sniffer.Ansible{Interval:1,Exe:"/data/kris/Cronus/Prometheus/plugin/ansible-sniff.py"}
        //an.Run()
        r,err:=sniffer.RunAnsibleCommand("/data/kris/Cronus/Prometheus/plugin/ansible-sniff.py","192.168.33.81,192.168.33.82")
        if err!=nil{
                fmt.Println(err.Error())
                return
        }
        servers,err:=sniffer.DecodeAnsibleReturn(r)
        if err!=nil{
                fmt.Println("json")
                fmt.Println(err.Error())
                return
        }
        fmt.Println(servers)
}
