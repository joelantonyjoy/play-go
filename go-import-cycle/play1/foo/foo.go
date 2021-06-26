package foo

import "play1/foobar/foo/bar"

var HELLO = "hello_from_foo"

type Foo1 interface {
	Method11()
}

type foo1 struct {
}

func (s *foo1) Method11() {
}

// ---

type Foo4 interface {
	Method41()
}

type foo4 struct {
	bar3 bar.Bar3 // IMP - foo4 can use Bar3 so "import cycle" problem is solved by moving interface definition to different package
}

func (s *foo4) Method41() {
	s.bar3.Method31()
}
