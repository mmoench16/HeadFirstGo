# Head First Go - Notes - Chapter 6

## Note 1

### Arrays

An array is a list of values of a particular type. 
Each item in an array is referred to as an array element.
An array holds a specific number of elements; no means are available to easily add more elements to an array.

## Note 2

### Slices

A slice is also a list of elements of a particular type, but unlike arrays, tools are available to add or remove elements.  

Slices don't hold any data themselves. A slice is merely a view into the elements of an underlying array.

## Bullet Points

+ The type for a slice is declared just like the type for an array variable, except the length is omitted:

*var mySclice []int*

+ For the most part, code for working with slices is identical to code that works with arrays. This includes: accessing elements, using zero values, passing slices to the *len* function, and *for...range* loops.
+ A **slice literal** looks just like an array literal, except the length is omitted:
*[]int{1, 7, 10}*
+ You can get a slice that contains elements *i* through *j* - 1 of an array or slice the **slice operators:** s[i:j]
+ The *os.Args* package variable contains a slice of strings with the command-line arguments the current program was run with.
+ A **variadic function** is one that can be called with a varying number of arguments.
+ To declare a variadic function, place an ellipsis (...) before the type of the last parameter in the function declaration. That parameter will then receive all the variadic arguments as a slice.
+ When calling a variadic function, you can use a slice in place of the variadic arguments by typing an ellipsis after the slice:

*in Range(1, 10, mySlice...)*