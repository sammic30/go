// Pointer is a variable that holds the memory address of another variable.

package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // & is to point
	log.Println("After func call myString is set to", myString)
}

func changeUsingPointer(s *string) { // * is pointer declaration
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
