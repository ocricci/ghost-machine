package main

import (
	"fmt"

	"example.com/ghost-machine/algo/binary"
)

func main() {
	fmt.Println("binary search: ", binary.Search([]int{1, 2, 3}, 1))
	fmt.Println("binary search rotation 1: ", binary.Rotationv1([]int{8, 9, 10, 2, 5, 6}))
	fmt.Println("binary search rotation 1: ", binary.Rotationv1([]int{2, 5, 6, 8, 9, 10}))
	fmt.Println("binary search rotation 1: ", binary.Rotationv1([]int{11, 12, 8, 9, 10}))
	fmt.Println("binary search rotation 1: ", binary.Rotationv1([]int{11, 12, 13, 15, 8, 9, 10}))

	fmt.Println("binary search rotation 2: ", binary.Rotationv2([]int{8, 9, 10, 2, 5, 6}))
	fmt.Println("binary search rotation 2: ", binary.Rotationv2([]int{2, 5, 6, 8, 9, 10}))
	fmt.Println("binary search rotation 2: ", binary.Rotationv2([]int{11, 12, 8, 9, 10}))
	fmt.Println("binary search rotation 2: ", binary.Rotationv2([]int{11, 12, 13, 15, 8, 9, 10}))
}
