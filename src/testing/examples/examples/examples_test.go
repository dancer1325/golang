package examples

// 1. example functions
func ExampleHello() {
	Hello()
	// Output: hello
}

// 1.2 error, because HelloReturned does NOT print anything
func ExampleHelloReturned() {
	HelloReturned()
	// Output: hello
}

// 1.3 MULTIPLE lines
func ExampleSalutations() {
	Salutations()
	// Output:
	// hello, and
	// goodbye
}

// 1.4 "Unordered output:"
func ExamplePerm() {
	Perm()
	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}

// 1.5 provide > 1 examples
func ExampleAdd() {
	Add(1, 2)
	// Output: 3
}

// suffix "_anotherThing"
func ExampleAdd_withNegative() {
	Add(1, -2)
	// Output: -1
}
