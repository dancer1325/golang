package main

import "fmt"

func main() {
	// 1. const identifier Type = expression				// Type				OPTIONAL
	const Pi float64 = 3.14159265358979323846

	// 2. expression NOT assignable -- to -- `Type`
	// const expressionNotAssignableToThatType float64 = "Alfred"		// uncomment to see the error

	// 3. const identifier = expression
	const zero = 0.0 // untyped floating-point constant

	// 4. const (
	//   ConstSpec1
	//   ConstSpec2
	//
	const (
		size int64 = 1024 // break line			-> ; 		NOT needed
		eof        = -1
	)

	// 5. const (ConstSpec1; ConstSpec2)
	const (
		sizeSame int64 = 1024
		eofSame        = -1
	)

	// 6. const identifier1, identifier2, ... = expression1, expression2, ...
	// type is omitted -> EACH constant's type == corresponding expressions's type
	const a, b, c = 3, 4, "foo"

	// 7. const identifier1, identifier2, ... Type = expression1, expression2, ...
	const u, v float32 = 0, 3 // Type specified -> ALL constants have the SAME type

	// 8. `Type` is omitted & expression values are untyped -> declared constants remain untyped
	const PiTypeOmittedAndUntyped = 3.14159
	fmt.Printf("PiTypeOmittedAndUntyped 's type %T & 's value %v\n", Pi, Pi) // TODO: what does "remain untyped" mean?

	// 9. parenthesized const declaration list's
	// 9.1 FIRST `ConstSpec` MUST have `ExpressionList`
	// 9.2 + `iota`
	const (
		Sunday = iota // FIRST `ConstSpec`			comment it to see the error
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Partyday
		numberOfDays // this constant is not exported
	)

}
