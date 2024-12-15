package main

type A struct{}

// Print1 prints number 1
func (A) Print1() {
	println("1")
}

// Print2 prints number 2
func (*A) Print2() {
	println("2")
}

// Print3 prints number 3
func (a A) Print3() {
	println("3")
}

// Print4 prints number 4
func (a *A) Print4() {
	println("4")
}

func main() {
	a := A{}
	a.Print1()
	a.Print2()
	a.Print3()
	a.Print4()
}
