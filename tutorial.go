package main

import (
	"fmt"
)

var i int = 12

func main() {
	a := []int{1, 2, 3}
	b := a[:]
	b[1] = 10
	fmt.Println(a)
	fmt.Println(b)
}
