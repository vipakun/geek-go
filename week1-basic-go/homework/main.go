package main

import "fmt"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, fmt.Errorf("out of boundaries at, %d", index)
	}
	res := src[index]

	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}

	src = src[:length-1]

	src = Shrink(src)
	return src, res, nil
}

func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCapacity(c, l)

	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

func calCapacity(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	}

	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}

	if c <= 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, true
}
