package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang/example/stringutil"
)

func main() {
	log.Println("Makefile Demo")

	http.Handle("/", http.HandlerFunc(helloHandler))
	http.Handle("/reverse", http.HandlerFunc(reverseHandler))
	http.ListenAndServe(":8080", nil)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello makefile\n"))
}

func reverseHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	reverseName := helloReverse(name)
	res := fmt.Sprintf(`%s%s`, reverseName, "\n")
	rw.Write([]byte(res))
}

func helloReverse(name string) string {
	return "hello " + stringutil.Reverse(name)
}
