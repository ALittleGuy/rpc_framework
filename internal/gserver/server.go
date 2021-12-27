package gserver

import (
	"fmt"
	"net"
	"rpc/internal/common"
	"rpc/internal/giface"
	"strconv"
)

type Server struct {
	name       string             // application name
	addr       string             // server addr
	port       uint32             // server port
	router     map[string]Service // service router
	dataPacker giface.IDataPack
}

func NewServer(name string, addr string, port uint32) *Server {
	return &Server{name: name, addr: addr, port: port, router: make(map[string]Service), dataPacker: NewDataPack()}
}

func (s *Server) RegisterService(service struct{}, name string) {
	if name == "" {
		name = service.
	}
	if _, ok := s.router[name]; ok {
		common.Assert(fmt.Sprintf("duplicate registration on '%s'", name))
	}
	s.router[name] = service
	return
}

func (s *Server) registerService(service struct{}) Service {

}

func (s *Server) RegisterFunction() {
	panic("implement me")
}

func (s *Server) Start() {
	fmt.Printf("Server %s listen on %s:%d start\n", s.name, s.addr, s.port)
	s.serve()
	select {}
}

func (s *Server) Stop() {

}

func (s *Server) serve() {
	socket := s.addr + ":" + strconv.Itoa(int(s.port))
	listener, err := net.Listen("tcp", socket)
	if err != nil {
		common.Assert(fmt.Sprintf("Server %s listen on %s:%d err\n", s.name, s.addr, s.port))
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("conn accept err", err)
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	for {
		select {
		default:
			msg, err := s.dataPacker.UnPack(conn)
			if err != nil {
				fmt.Print("request read fail ", err)
			}
			request

		}
	}
}
