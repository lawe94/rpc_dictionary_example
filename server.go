package main

import (
	"net"
	"net/rpc"
	"net/http"
	"log"
	"examples/rpc/server"
)

func main(){
	store := server.NewStore()
	rpc.Register(store)
	rpc.HandleHTTP()
	l, e := net.Listen("unix", "/tmp/rpc.sock")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}