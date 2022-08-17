package main

import (
	"fmt"
	"strconv"
)

var i int = 12

func main() {
	i := 120
	j := strconv.Itoa(i)
	k := string(i)
	fmt.Println(j)
	fmt.Println(k)
}
