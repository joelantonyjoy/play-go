package main

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

type Employee struct {
	Name string
	Age  int
}

func main() {
	log.SetFlags(log.Lmicroseconds)

	log.Printf("sending ..\n")
	conn, err := net.Dial("tcp", "localhost:8050")
	if err != nil {
		log.Fatalf("Dialing Error, %v", err)
	}

	enc := gob.NewEncoder(conn)
	err = enc.Encode(Employee{Name: "nurali", Age: 35})
	if err != nil {
		log.Fatalf("Encoding failed, %v", err)
	}
	log.Printf("sent\n")

	time.Sleep(2 * time.Second)
	err = conn.Close()
	if err != nil {
		log.Fatalf("Close failed, %v", err)
	}
	log.Printf("done\n")
}
