package main

import (
	"rpc/internal/gserver"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	server := gserver.NewServer("0.0.0.0" , 8080)
	server.Start()

}
