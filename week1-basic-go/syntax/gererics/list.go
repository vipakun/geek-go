package gererics

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkedListV1[T any] struct {
	head *nodeV1[T]
}

type nodeV1[T any] struct {
	data T
}

func (l LinkedListV1[T]) Add(index int, val T) {

}

func UseList() {
	l := &LinkedListV1[int]{}
	l.Add(1, 123)

	types.Link
}
