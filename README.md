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
<br>
# Comparison Operators
```
== equal
!= not equal
< less than
<= less than or equal to
> greater than
>= greater than or equal to
```
# Arithmetic Operators
```
<br>
+ addition
- subtraction
* multiplication
/ division
% modulus
++ increment
-- decrement
```
# Logical Operators
```
<br>
&& AND - // returns true, if both statements are true and returns false, if either of the statement is false.
|| OR - // returns true, if one of the statement is true and returs false, when both statements are false.
! NOT - // reverses the results, returns false if the expression evaluates to true and vice versa.
```
# Assignment Operators
```
<br>
= assign
+= add and assign    // x+= y means x = x + y (assigns left operand with the addition result)
-= sustract and assign
*= multiply and assign
/= divide and assign quotient // x/=y means x = x / y
%= divide and assign modulus
```
# Bitwise Operators
<br>
```
& AND
| OR
^ XOR - result of XOR is 1 if the two bits are opposite, example 0 and 1
>> Right Shift
<< Left Shift - shifts all bits towards left by a certain number of specified bits. The bit position that have been vacated by the left shift operator are filled with 0
```

