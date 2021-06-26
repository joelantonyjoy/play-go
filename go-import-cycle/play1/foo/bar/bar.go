package bar

import (
	"play1/foobar/foo"
)

var HELLO = "hello_from_bar"

type Bar2 interface {
	Method21()
}

type bar2 struct {
	foo1 foo.Foo1
}

func (f *bar2) Method21() {
	f.foo1.Method11()
}

// ---

type Bar3 interface {
	Method31()
}

type bar3 struct {
}

func (f *bar3) Method31() {
}
