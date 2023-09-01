Syntax:
    if condition {
        // execute when condition is true
    }

package main

import (
	"fmt"
)

func main() {
	age := 17
	if age >= 18 {
		fmt.Println("You are eligible for voting:", age)
	} else {
		fmt.Println("You are NOT eligible for voting:", age)
	}
}

// Another Example

package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanf("%s", &age)
	if age >= 18 {
		fmt.Println("You are eligible for voting:", age)
	} else {
		fmt.Println("You are NOT eligible for voting:", age)
	}
}
