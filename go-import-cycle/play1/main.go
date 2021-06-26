package main

import (
	"fmt"

	"play1/foo"
	"play1/foo/bar"
)

func main() {
	fmt.Println("play1")

	fmt.Println(foo.HELLO)
	fmt.Println(bar.HELLO)
}
