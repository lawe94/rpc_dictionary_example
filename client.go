package main

import (
	"net/rpc"
	"examples/rpc/server"
	"fmt"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 3{
		fmt.Println("Usage: 'get key' | 'set key value")
		os.Exit(1)
	}
	option := os.Args[1]
	switch option{
		case "get":
			key := os.Args[2]
			get(key)
		case "set":
			key := os.Args[2]
			value := os.Args[3]
			set(key,value)
		default:
			fmt.Println("unknown option")
			os.Exit(1)
	}
}

func get(key string){
	client, err := rpc.DialHTTP("unix","/tmp/rpc.sock")
	returnValue := new(server.ReturnValue)
	err = client.Call("KeyValueStore.Get", key, returnValue)
	if err != nil {
		log.Fatal("KeyValue error Get:", err)
	}
	fmt.Println(returnValue)
}

func set(key, value string){
	entry := server.Pair{key,value}
	client, err := rpc.DialHTTP("unix","/tmp/rpc.sock")
	returnValue := new(bool)
	err = client.Call("KeyValueStore.Set", entry, returnValue)
	if err != nil {
		log.Fatal("KeyValue error Get:", err)
	}
	fmt.Printf("set %s to %s", key, value)
}