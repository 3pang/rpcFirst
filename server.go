package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	serverApi "rpcFirst/server"
)

func main1() {
	log.Println("rpc 服务端启动")
	api := new(serverApi.API)
	api2 := new(serverApi.API)
	err := rpc.Register(api)
	err2 := rpc.RegisterName("api2", api2)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	if err2 != nil {
		log.Fatal("error registering API2", err2)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("initial database: ", database)
	// a := Item{"first", "a test item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("second database: ", database)

	// DeleteItem(b)
	// fmt.Println("third database: ", database)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("fourth database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)

}
