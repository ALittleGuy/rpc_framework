package giface

import "rpc/internal/gserver"

type Function func() error

type IServer interface {
	RegisterService(service gserver.Service , name string) // register a function for rpc
	RegisterFunction()
	Start()                                      // start server
	Stop()                                       // stop server and close all chan
}
