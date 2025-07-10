package main

import "fmt"

func main() {
	// 1. use cases
	// 1.1 variable declaration
	var name string // ""
	fmt.Println(name)
	var age int // 0
	fmt.Println(age)
	var active bool // false
	fmt.Println(active)
	var ptr *int // nil
	fmt.Println(ptr)

	// 1.2 call of `new`
	anotherptr := new(int) // *ptr = 0
	fmt.Println(anotherptr)
	str := new(string) // *str = ""
	fmt.Println(str)

	// 1.3 create NEW value -- through -- composite literals
	type Person struct {
		Name string
		Age  int
	}

	p := Person{} // {Name: "", Age: 0}
	fmt.Println(p)
	arr := [3]int{} // [0, 0, 0]
	fmt.Println(arr)

	// 1.4 create NEW value -- through -- call of `cmake`
	slice := make([]int, 3) // [0, 0, 0]
	fmt.Println(slice)
	m := make(map[string]int)
	fmt.Println(m)
	ch := make(chan int)
	fmt.Println(ch)

	// 2. zero-value applies recursively
	type Address struct {
		Street string
		Number int
	}
	type Company struct {
		Name    string
		Address Address
		Active  bool
	}
	var company Company

	fmt.Printf("%+v\n", company)

	// 3. zero-values can ALSO be specified EXPLICITLY
	var a int
	fmt.Println(a)
	var b int = 0
	fmt.Println(b)
}
