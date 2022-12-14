package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Alice"
	want := "Hello Alice"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
