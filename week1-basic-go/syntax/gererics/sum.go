package gererics

import "io"

func Sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res = res + v
	}
	return res
}

type Integer int

type Number interface {
	// int | uint | int32
	~int
	// ~int: int and  all the int similar types, like int32, or custom type Integer
}

func UserSum() {
	res := Sum[int]([]int{12, 123})

	resV1 := Sum[Integer]([]Integer{12, 123})
}

func Closable[T io.Closer]() {
	var t T
	t.Closer()
}

func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t < vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 1; i < len(vals); i++ {
		if t > vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Find[T Any](vals []T, filter func(t T) bool) T {
	for _, v := range vals {
		if filter(v) {
			return v
		}
	}
	var t T // cannot find, return 0 (default)
	return t
}

func Insert[T any](idx int, val T, vals []T) []T {
	if idx < 0 || idx >= len(vals) {
		panic("idx is not valid")
	}
	vals = append(vals, val)
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 > 0 {
			vals[i] = vals[i-1]
		}
	}
	vals[idx] = val
	return vals
}
