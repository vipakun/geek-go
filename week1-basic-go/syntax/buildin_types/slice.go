package main

import "fmt"

func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("s1: %v, len: %d, cap: %d", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4) // create a slice with length 3, capicity 4 slice
	fmt.Printf("s2: %v, len: %d, cap: %d", s2, len(s2), cap(s2))

	s2 = append(s2, 7)
	fmt.Printf("s2: %v, len: %d, cap: %d", s2, len(s2), cap(s2))
}
