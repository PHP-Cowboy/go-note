package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (r *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	_ = rpc.RegisterName("HelloService", &HelloService{})

	conn, err := listener.Accept()

	if err != nil {
		panic(err)
	}

	rpc.ServeConn(conn)

}
