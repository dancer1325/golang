package main // package block's start,  file block's start
import "fmt"

func main() {
	universeBlock()

	ifBlock()

	forBlock()

	switchBlock()

	selectBlock()
}

func selectBlock() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	select {
	case val := <-ch1:
		// case val's block
		result := val * 2
		fmt.Println(result)
		// 'val' & 'result' ONLY exist here
	case msg := <-ch2:
		// case msg's block
		result := "Got: " + msg
		fmt.Println(result)
		// 'msg' & 'result' ONLY exist here
	}
}

func switchBlock() {
	x := 2
	switch x {
	case 1:
		// case1 block
		msg := "One"
		fmt.Println(msg)
		// 'msg' ONLY exist here
	case 2:
		// case2 block
		msg := "Two" // != 'msg' | case1
		fmt.Println(msg)
		// 'msg' ONLY exist here
	}
}

func forBlock() {
	for i := 0; i < 5; i++ { // ← 'i' ONLY exists | this for block
		fmt.Println(i) // 'i' accessible here
	}
	// fmt.Println(i)         // ❌ Error: 'i' NOT exist here
}

func ifBlock() {
	x := 10

	if y := x * 2; y > 15 { // 'y' ONLY exists | this if block
		fmt.Println(y) // 'y' accessible here
	}
	// fmt.Println(y)        // ❌ Error: 'y' NOT exist here
}

func universeBlock() {
	// universe block's basic types
	var x = 42
	var y = "hello"
	var z = true

	// universe block's built-in functions
	slice := make([]int, 5)
	length := len(slice)
	capacity := cap(slice)
}

// if there is NO more files | this package -> package block's end,  file block's end
