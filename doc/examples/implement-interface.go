package main

import "fmt"

// 1. `T` NOT interface && == element of `I`'s type set
// `I`'s type set
type I interface {
	~int | ~string // type set: int y string
	//int | string 	// TODO: what does `~` mean? why is it necessary?
}

type T int // int == `I`'s type set

// generic function / type parameter == `I`
func process[V I](value V) {
	fmt.Println(value)
}

// 2. S == interface & `S`'s type set == subset of the `R`'s type set
// R's type set > S's type set
type R interface {
	~int | ~string | ~float64
}

// S's type set included | R's type set
type S interface {
	~int | ~string
}

// generic function / type parameter == R
func processR[V R](value V) {
	fmt.Printf("processR: %v\n", value)
}

// generic function / type parameter == S
func processS[V S](value V) {
	fmt.Printf("processS: %v\n", value)
	processR(value) // S implements R
}

type OnlyInt int
type OnlyString string

// 3. value of type `M` implements `N`
type N interface {
	String() string
	Value() int
}

type M struct {
	data int
}

func (m M) String() string {
	return fmt.Sprintf("T{%d}", m.data)
}

func (m M) Value() int {
	return m.data
}

func useInterface(n N) {
	fmt.Printf("String: %s, Value: %d\n", n.String(), n.Value())
}

func main() {
	// 1.
	var t T = 42
	process(t) // T can be passed == T implements I
	// NOT ALLOWED -- TODO: why?
	//var i I = t
	//fmt.Println(i)

	// 2.
	processR(OnlyInt(10)) // MyInt implements S & S implements R

	// 3.
	m := M{data: 42}
	useInterface(m)
}
