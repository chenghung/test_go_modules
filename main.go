package main

import (
	"github.com/kr/pretty"
)

type test struct {
	v interface{}
	s string
}

func main() {
	t := test{
		v: 5,
		s: "test",
	}
	pretty.Log(t)
}
