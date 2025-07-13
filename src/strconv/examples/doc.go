package main

import (
	"fmt"
	"strconv"
)

func main() {
	numericConversions()
	stringConversions()
}

func numericConversions() {
	// 1. Atoi
	i, _ := strconv.Atoi("-42") // string -- to -- int
	fmt.Printf("Atoi: int converted %v with type %T\n", i, i)

	// 2. Itoa
	s := strconv.Itoa(-42) // int -- to -- string
	fmt.Printf("Itoa: string converted %v with type %T\n", s, s)
}

func stringConversions() {
	// 1. Quote
	q := strconv.Quote("Hello, 世界")
	fmt.Println("1. Quote ", q)

	// 2. QuoteToASCII
	r := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println("2. QuoteToASCII ", r) // non-ASCII Unicodes are escaped -- with -- `\u`

	// 3. QuoteRune
	rr := strconv.QuoteRune('世')
	fmt.Println("3. QuoteRune ", rr) // TODO: Why quoted with '  != BEFORE with "

	// 4.  QuoteRuneToASCII
	rrr := strconv.QuoteRuneToASCII('世')
	fmt.Println("4.  QuoteRuneToASCII ", rrr) // TODO: Why quoted with '  != BEFORE with "

	// 5. Unquote
	uq, err := strconv.Unquote(`"Hello, 世界"`)
	fmt.Println("5. Unquote ", uq, err)

	// 6. UnquoteChar
	uc, multibye, tail, err := strconv.UnquoteChar(`\u4e16`, 0)
	fmt.Println("6. UnquoteChar ", uc, multibye, tail, err)
}
