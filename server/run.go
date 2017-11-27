package server

import (
	"net/http"
	"fmt"
	"github.com/VinkDong/jsond/common"
)

func StartServer(address string) {
	fmt.Printf("starting server at %s%s%s\n", common.CLR_Y, address, common.CLR_N)
	common.Server.Err = http.ListenAndServe(address, nil)
}