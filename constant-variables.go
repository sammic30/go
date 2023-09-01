// Contstant value cannot be changed
// We have to define and declare constant at the same time
// We cannot use short hand declaration with constant, example: const name := "Maggie" -- This will trow error

package main

import (
	"fmt"
	"strconv"
)

func main() {
	const name = "Harry Potter"
	const is_muggle = false
	const age = 12

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n", age, age)
}

------

package main

import (
	"fmt"
)

func main() {
	const name = "Harry Potter"
	name = "Blah" // This will throw error, b'coz we tried to change the constant variable's value
	const is_muggle = false
	const age = 12

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n", age, age)
}

// Above code is throwing "./main.go:9:2: cannot assign to name (untyped string constant "Harry Potter")" error, b'coz we tried to change the constant variable's value

// constant usage example

package main

import (
	"fmt"
)

const PI float64 = 3.14 // global constant

func main() {
	radius := 5.0
	var area float64
	area = PI * radius * radius

	fmt.Println("Area of the Circle is: ", area)
}
