package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	err := rpc.Register(TestService{})
	if err != nil {
		fmt.Print("err ", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":8080", nil)

}

type TestService struct{}

func (t *TestService) Add(a, b int) error {
	return a + b
}
