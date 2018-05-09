package main

import "testing"

func TestHelloReverese(t *testing.T) {
	expected := "hello lameh"
	name := "hemal"
	actual := helloReverse(name)
	if actual != expected {
		t.Errorf("result was incorrect, want:%s, got%s", expected, actual)
	}
}
