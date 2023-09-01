/*
Converting Variables from one Data Type to Another (Type Casting)
Itoa()
    - converts integer to string.
    - returns one value - string formed with the given integer.

Atoi()
    - converts string to integer.
    - returns two values - the corresponding integer, error (if any).
*/
// Converting int into a floating number
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
}

// Convert integer to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42
	newstring := strconv.Itoa(num) // convert int to string
	fmt.Printf("%q\n", newstring)
}

// string to integer convertion

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "420"
	i, err := strconv.Atoi(num) // store return value... i have the value of num now
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)
}
