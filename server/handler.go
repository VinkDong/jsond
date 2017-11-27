package server

import "net/http"

func HomeHandler(resp http.ResponseWriter, req *http.Request)  {
	resp.Write([]byte("ok"))
}