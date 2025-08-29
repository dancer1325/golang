package main

import "fmt"

func anotherFunctionTwo() int { return 42 }
func tupleAssignmentFunction() (int, int) { return 42, 23 }

func main() {
	var x int // declare the variable
	x = 1     // assignment
	fmt.Print("assignment ", x)

	// 2. address operator == *
	var val int // declare the variable
	p := &val
	*p = anotherFunctionTwo()

	// 3. map index expression
	a := make(map[int]int)
	i := 0
	a[i] = 23

	// 3. operand can be parenthesized
	ch := make(chan int, 1)
	ch <- 100
	var k int
	(k) = <-ch // same as: k = <-ch

	// 4. assignment operation
	x += 5
	x *= 2

	// 5. tuple assignment
	// 5.1 right hand operand == single multi-valued expression
	b, c = tupleAssignmentFunction()		// declared PREVIOUSLY | another package's file
	// 5.2 number of operands | left == number of expressions | right
	d, g, h = '一', '二', "三"

	// 6. _		== blank identifier
	var s			// declare BEFORE assign
	_, s = 1, 2		// ignore the first value
	l, _ := 1, 2	// ignore the second value
}
