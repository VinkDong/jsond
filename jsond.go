package main

import (
	"flag"
	"github.com/VinkDong/jsond/common"
	"os"
	"fmt"
	"github.com/VinkDong/jsond/server"
	"github.com/VinkDong/asset-alarm/log"
	"time"
)

var (
	showHelp = flag.Bool("help", false, "show help")
	address  = flag.String("addr", "", "service address")
	port     = flag.Int("port", 8081, "service port")
	config   = flag.String("conf", "", "Config file")
)

func main() {
	flag.Parse()

	if *showHelp {
		common.Help()
		os.Exit(0)
	}

	listenAddress := fmt.Sprintf("%s:%d", *address, *port)
	go server.StartServer(listenAddress)
	for {
		time.Sleep(time.Second)
		if common.Server.Err != nil{
			log.Errorf("start server at %s failed",listenAddress)
			os.Exit(128)
		}
	}
}
