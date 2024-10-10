package main

import "fmt"

type Fish struct {
}

func (f Fish) Swim() {

}

func (f FakeFish) FakeSwim() {

}

type FakeFish Fish

// FakeFish is a new type

func UseFish() {
	f1 := Fish{}

	f1.Swim()

	f2 := FakeFish{}
	f2.FakeSwim()

	f3 := FakeFish(f1)
	f1 := Fish(f3)
	fmt.Println(f1, f3)

}
