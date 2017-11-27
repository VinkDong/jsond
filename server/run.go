package server

import (
	"net/http"
	"fmt"
	"github.com/VinkDong/jsond/common"
)

func InitHandler()  {
	http.HandleFunc("/",HomeHandler)
}

func StartServer(address string) {
	InitHandler()
	fmt.Printf("starting server at %s%s%s\n", common.CLR_Y, address, common.CLR_N)
	common.Server.Err = http.ListenAndServe(address, nil)
}