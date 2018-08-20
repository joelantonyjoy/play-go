package main

import (
	"encoding/gob"
	"log"
	"net"
)

type Emp struct {
	Name string
	Age  *int
}

func main() {
	log.SetFlags(log.Lmicroseconds)

	log.Printf("starting ..\n")
	addr, err := net.ResolveTCPAddr("tcp4", ":8050")
	if err != nil {
		log.Fatalf("Resolve failed, %v", err)
	}
	lisntener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Listen failed, %v", err)
	}
	log.Printf("listening ..\n")
	conn, err := lisntener.Accept()
	if err != nil {
		log.Fatalf("Conn failed, %v", err)
	}

	dec := gob.NewDecoder(conn)
	var e1 Emp
	log.Printf("receving ..\n")
	err = dec.Decode(&e1)
	if err != nil {
		log.Fatalf("decode failed, %v", err)
	}
	log.Printf("Got data, Emp:%v, name:%s, age:%d", e1, e1.Name, *e1.Age)
}
