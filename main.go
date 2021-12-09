package main

import (
	"fmt"
	"slice/search"
)

func main() {
	a := []int{1, 2, 3, 8, 9, 3, 2, 1}
	fmt.Println(*search.MaxValueSequenceAndSequenceInvers(a))
}
