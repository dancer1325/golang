package abs

import "math"

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}
