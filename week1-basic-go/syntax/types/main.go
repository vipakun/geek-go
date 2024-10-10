package main

import "fmt"

func main() {
	u1 := &User{}
	println(u1)
	u1 = new(User)
	println(u1)

	u2 := User{}
	fmt.Printf("%v \n", u2)

	u2.Name = "Xiangkun Yang"
	println(u2.Name)

	var u3 User
	fmt.Printf("%+v \n", u3)

	var u4 *User
	println(u4)

	u5 := User{Name: "XY"}
	fmt.Printf("%+v \n", u5)

	ChangeUser()
	UseFish()
}

func UseList() {
	// l1 := LinkedList{}
	// l1Ptr := &l1
	// var l2 = *l1Ptr
	// fmt.Printf("%+v \n", l2)
	//
	// var l3Ptr *LinkedList
	// println(l3Ptr)
}
