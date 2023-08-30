# go

# Comments:
<br>
// - Single line comment:
<br>
Multi line comment:
<br>
/*
adasd
kjhsdf
*/
<br>
#-------------------------

# Data Types:
<br>
String
    <br>
    "abd"
    <br>
    "ABD kkjf Blah"

Numbers
    <br>
    Integer
    <br>
    Float

Boolean:
    <br>
    True
    <br>
    False

Array & Slices:
    <br>
    [1,2,4,6,7,34]
    <br>
    ["foo","bar","blah"]
    <br>
    [7.0,2.4,5.6]

#Arrays are a fixed-length data structure, meaning that their size is known at compile time.
<br>
#Slices, on the other hand, are a variable-length data structure, meaning that their size can change at runtime.

Maps // Is a collection of key value pairs:
    <br>
    "x" -> 30
    <br>
    1 -> 100

#-------------------------
<br>

# Variables:

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
<br>
# Variable Scope:

Blocks
{ // outer block

    {
        // inner block
    }

}

Inner blocks can access variables declared within outer blocks
Outer blocks cannot access variables declared within inner blocks

#------------------------

