package main

// Import packages -- to use -- their exported identifiers
import (
	"fmt"       // Standard package
	"math"      // Standard package for math functions
	"math/rand" // Subpackage of math
	"time"      // Package for time functions
)

func main() {
	// fmt.Println 			fmt == PackageName, 		Println == exported identifier
	fmt.Println("Examples of qualified identifiers:")

	// math.Pi 				math == PackageName,		Pi == exported identifier (constant)
	fmt.Println("Value of Pi:", math.Pi)

	// time.Now 			time == PackageName, 		Now == exported identifier (function)
	now := time.Now()
	fmt.Println("Current time:", now)

	// rand.Intn 			rand == PackageName, 		Intn == exported identifier (function)
	randomNumber := rand.Intn(100)
	fmt.Println("Random number:", randomNumber)
}
