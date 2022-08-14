package main

import "fmt"

func main() {
	var group rune
	group = 'b'
	var code byte = 12
	var num float32 = 4.5
	var is_trusted = false
	var phone int8 = 28

	fmt.Println("Hello world!", group, code, num, is_trusted, phone)

	var name string = "homayoon"
	var age int16 = 23
	fmt.Println(name, age)
	fmt.Printf("%T", name)
	fmt.Printf("%T", age)
}
