package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	conn, err := listener.Accept()

	err = rpc.RegisterName("HelloService", &HelloService{})

	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
}
