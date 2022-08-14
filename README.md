# Golang Tutorial

## 1. Getting started

- After installing the Go package from the `go.dev` website, you can initialize your first `.go` file very simply. Go is a compiled language and an executable file will be created after compilation.
- As a simple first steps, first let's assing `package main` which at the moment I'm not sure what is does! But I guess it somehow points out that this .go file is going to be executed as a main file, not a side module or something. # TODO: Find this out!
- Second, `import` the `fmt` module with looks like to be a printing module or something.
- Then we have to create a `function` called `main`. Something that I just found out is that if the `package main` is not on top of the file, you cannot directly use `go run {file}` in the shell. It says that it's not a main package. and this function should also be called `main()` specifically. This might have somekind of correlation with the package type/name.
- Inside the funtion you can write `fmt.Println("Hello World!")` in order to print out this line, `''` is for `characters` like: `'a'` and `""` can be used to lenghtier (!) strings. at the end your program should look like this:
- You can also `build` your go file to create the mentioned `executable`. This file doesn't need Go to be installed on the target machine and can be run almost anywhere.

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

```shell
go build tutorial.go
=>>> Results in a binary looking (bytecode maybe?) file.
go run tutorial.go
=>>> Hello world!
```

---

## 2. Variables and Data Types

- Go is a staticly typed language and because of that, along with the declaration of each variable one should also declare it's type. Something like:

```go
func main() {
    var name string = "homayoon"
    var age int16
    age = 23
    fmt.Println(name, age)
}
```

- Note that you can declare a variable and it's type but not necessarily assign any value to it right at the moment.
- There are different types of variables in the Go language which are listed below:

```go
uint8 uint16 uint32 uint64
int8 int16 int32 int64
float32 float64
complex64 complex128
string
bool
```

- `uint8` is also called `byte`
- `int32` is also called `rune`

- If you're puting anything in Go, let it be a single line or a declaration or a variable, you `HAVE TO` use it or it will be get deleted automatcially!

- consider the size boundries of your variables before declaring them, `overflowing` is something that might happen for your variables.
- `Characters --> 'a'` are not meant to be considered `strings`, but they are in fact `int32 (rune)` and when printing them out (or maybe adding up to other chars/numbers) they appear to be an integer.

- Running the code below would result in the variables beings printed in the same line separated by a space.

```go
package main

import "fmt"

func main() {
	var name string = "homayoon"
	var group rune
	group = 'b'
	var age int8 = 23
	var code byte = 12
	var num float32 = 4.5
	var is_trusted = false
	var phone int8 = 28

	fmt.Println("Hello world!", name, group, age, code, num, is_trusted, phone)
}
```

```shell
go run tutorial.go
=>>> Hello world! homayoon 98 23 12 4.5 false 28
```

---

## 3. Implicit and Explicit definition. Default values and types

- You can define a variable in different ways, some of them are called `implicit` and others are `explicit`, which in the essence means if you are specifying the variable type yourself or not.

```go
var name string = "homayoon" // explicit
var name = "homayoon" // implicit
var age = 12 // implicit, might be an specific type of int or unit
// that you actually don't want.
name := "homayon" // still implicit
```

- the `:=` is called `Assignment Expression` or `Walrus operator` and does the exact same thing as the line above it.
- the `implicit` type of variable declaration forces Go itself to determine a type for your variable. THIS MIGHT BE PROBLEMATIC IN SOME CASES.
- the Walrus operator can only be used `inside of a function`. because outside a function, every line has to be started with a keyword like `func` or `var`.
- variables doesn't have to be given a value right at the beginning. you can assign them a suitable value later, but until then they are going to hold a `default value`.

```go
var name int // default to 0
var my_bool bool // defaults to false
```

- You can not `declare a variable twice` AT LEAST IN A SINGLE FUNCTION!
- You can print `type of a variable` by using `fmt.Printf("%T", your_var)`.

```go
var name string = "homayoon"
var age int16 = 23
fmt.Println(name, age)
fmt.Prinf("%T", name) // string
fmt.Print("%T", age) // int16
```

---
