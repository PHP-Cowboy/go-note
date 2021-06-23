package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	var replay string
	err = client.Call("HelloService.Hello", "cowboy", &replay)

	if err != nil {
		panic(err)
	}
	fmt.Println(replay)
}
