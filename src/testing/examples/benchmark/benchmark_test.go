package benchmark

import (
	"fmt"
	"math/rand"
	"testing"
)

// 1. syntax & naming
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ { //	target code is run 		b.N times / adjusted TILL benchmark function lasts long enough
		rand.Int()
		fmt.Print("BenchmarkRandInt i: ", i)
	}
}

// 2. reset timer
type Big struct {
	data []int
}

func NewBig() *Big {
	return &Big{
		data: make([]int, 1000), // Setup costoso
	}
}

// method to benchmark
func (b *Big) Len() int {
	return len(b.data)
}

// reset timer			Reason: benchmark needs expensive setup BEFORE running
func BenchmarkBigLen(b *testing.B) {
	big := NewBig()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		big.Len()
	}
}

// 3. run parallel
func BenchmarkRandIntParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Int()
		}
	})
}
