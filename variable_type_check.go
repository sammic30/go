package main

import "fmt"

func main() {
	grades := 42
	message := "Hello, People"
	isCheck := true
	amount := 54.66

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
}

// %T is to check the type

// OR

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Damini"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.02))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}
