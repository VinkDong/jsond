package common

import "fmt"

const VERSION  = "v0.1"

// color
const CLR_0 = "\x1b[30;1m"
const CLR_N = "\x1b[0m"
const CLR_R = "\x1b[31;1m"
const CLR_G = "\x1b[32;1m"
const CLR_Y = "\x1b[33;1m"
const CLR_B = "\x1b[34;1m"
const CLR_M = "\x1b[35;1m"
const CLR_C = "\x1b[36;1m"
const CLR_W = "\x1b[37;1m"

func Help()  {
	fmt.Printf(`%sJsonD %sJsond wish to make more flexible using for exists json api, we will persistently update it. The propose use like grafana simpleJson datasource backend etc.
--conf        set a config file,must be yaml
--addr        set listen address
--port        set listen port
--help        show help information
--version     show version
Need more refer at  %shttps://github.com/VinkDong/JsonD%s
`, CLR_Y, CLR_N, CLR_C, CLR_N)
}