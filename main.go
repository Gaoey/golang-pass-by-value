package main

import "fmt"

type Test struct {
	X string
}

func New(e string) Test {
	return Test{X: e}
}

func (t Test) Hello(e string) {
	newstr := t.X + e
	fmt.Printf("%s\n", newstr)
}

func (t Test) setX(x string) {
	t.X = x
}

func (t Test) printX() {
	fmt.Printf("%s\n", t.X)
}

func main() {
	t := New("Helloy")
	t.Hello("yo")
	t.Hello("yoooo")

	t.setX("eee") // cannot set
	t.printX()    // print Hello
}
