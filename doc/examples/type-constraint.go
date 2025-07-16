package main

// 1. define ALLOWED type arguments
type Numeric interface {
	int | float64 | string // ONLY ALLOWED types
}

// TODO:

// 2. CONTROLS OPERATIONS SUPPORTED
type Comparable interface {
	int | string
}

func main() {

}
