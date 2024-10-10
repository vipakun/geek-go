package main

import (
	"variables/demo"
)

var (
	a1 int     = 123
	a2 float64 = 12.3
)

func main() {
	// var a int = 10
	println(demo.Global)

}

const (
	Status0 = iota
	Status1
)
