package funcs

func Closure(name string) func() string {
	return func() string {
		return "Hello " + name
	}
}

func ClosureInvoke() {
	c := Closure("XY")
	println(c())
}
