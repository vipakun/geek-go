package main

func Defer() {
	defer func() {
		println("first defer")
	}()

	defer func() {
		println("second defer")
	}()
}

func DeferClosure() {
	j := 0
	defer func() {
		println(j)
	}()
	j = 1
}

func DeferClosureV1() {
	j := 0

	defer func(j int) {
		println(j)
	}(j)
	j = 1
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
