package main

import (
	"net"
	"net/rpc"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	go closeSocket()
	http.Serve(l, nil)
}

func closeSocket(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	syscall.Unlink("/tmp/rpc.sock")
	os.Exit(0)
}