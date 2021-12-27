package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"rpc/internal/gserver"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	Ans string
}

func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer   // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	gob.Register(P{})
	gob.Register(Q{})
	// Encode (send) the value.
	req := gserver.Request{Param: P{Name: "asd"} , Reply: Q{Ans: "asdasd"}}
	err := enc.Encode(req)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q gserver.Request
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Print( q.Param)
	fmt.Print()
}
