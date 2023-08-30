# go

Comments:
// - Single line comment

Multi line comment
/*
adasd
kjhsdf
*/

#-------------------------

String
    "abd
    "ABD kkjf Blah"

Numbers
    Integer
    Float

Boolean
    True
    False

Array & Slices 
    [1,2,4,6,7,34]
    ["foo","bar","blah"]
    [7.0,2.4,5.6]

#Arrays are a fixed-length data structure, meaning that their size is known at compile time.
#Slices, on the other hand, are a variable-length data structure, meaning that their size can change at runtime.

Maps // Is a collection of key value pairs
    "x" -> 30
    1 -> 100

#-------------------------

Variables:

Declaring a variable

Syntax
    var <variable name> <variable type> = <value>

    var s string = "Hello, World!"
    var i int = 100
    var b bool = false
    var f float64 = 77.90

Short hand way
    var s,t string = "foo", "bar"

    var (
    s sting = "foo, bar"
    i int = 100
    )

Short variable declaration
    s := "foo, bar"

#------------------------

Variable Scope:

Blocks
{ // outer block

    {
        // inner block
    }

}

Inner blocks can access variables declared within outer blocks
Outer blocks cannot access variables declared within inner blocks

#------------------------
