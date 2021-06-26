package foo

import "play0/foo/bar"

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
	bar3 bar.Bar3 // IMP - foo4 can NOT use Bar3 (here Bar3 belongs to package "play0/foo/bar" while "play0/foo/bar" already import "play0/foo" which creates "import cycle" problem)
}

func (s *foo4) Method41() {
	s.bar3.Method31()
}
