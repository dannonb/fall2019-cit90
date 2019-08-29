package main 

import "testing"

func TestFoo(t *testing.T) {
	want:="hello"
	if got := Foo(); got != want {
		t.Errorf("foo() = %v, but got %v", got, want)
	}
}