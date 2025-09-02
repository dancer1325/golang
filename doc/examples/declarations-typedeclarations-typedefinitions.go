package main

import "fmt"

// 1. type (TypeSpec;TypeSpec; ...)
type (
	// identifier Type	-- identifier TypeLit  -- identifier StructType
	PointAnotherTwo struct{ x, y float64 }
	// identifier Type  --
	polarAnotherTwo int
)

//  2. type (
//     TypeSpec
//     TypeSpec
//     ...
//     )
type (
	PointAnother struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar        int                    // polar and Point denote different types
)

// 3. type TypeSpec	 --  type TypeDef  --  type identifier Type  --  type identifier TypeLit  --  type identifier StructType
type TreeNode struct {
	left, right *TreeNode
	value       any
}

// 4. type TypeSpec	 --  type TypeDef  --  type identifier Type  --  type identifier TypeLit  --  type identifier InterfaceType
type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

// 4.1 implement type interface
type SimpleBlock struct {
	size int
}

func (b SimpleBlock) BlockSize() int {
	return b.size
}

func (b SimpleBlock) Encrypt(src, dst []byte) {
	copy(dst, src)
}

func (b SimpleBlock) Decrypt(src, dst []byte) {
	copy(dst, src)
}

// TODO:

func main() {
	pointAnother := PointAnother{1, 2}
	fmt.Printf("1. pointAnother %v\n", pointAnother)

	polar := 2
	fmt.Printf("2. polar %v\n", polar)

	tree := TreeNode{nil, nil, 1}
	fmt.Printf("3. tree %v\n", tree)

	// define -- as -- Block
	var block Block = SimpleBlock{size: 16}
	fmt.Printf("4. block's type: %T\n", block)
}
