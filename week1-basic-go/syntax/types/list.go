package main

type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}

type LinkedList struct {
	head node
}

func (l *LinkedList) Add(index int, val any) {
	// implement

}

type node struct {
}
