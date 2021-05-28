# Head First Go - Notes - Chapter 5

## Note 1

### Arrays

An array is a list of values of a particular type. 
Each item in an array is referred to as an array element.
An array holds a specific number of elements; no means are available to easily add more elements to an array.

## Bullet Points

+ To declare an array variable, include the array length in square brackets and the type of elements it will hold:

*var myArray [3]int*

+ To assign or access an element of an array, provide its index in square brackets. Indexes start at 0, so the first element of *myArray* is *myArray[0]*.
+ As with variables, the default value for all array elements is the zero value for that element's type.
+ You can set element values at the time an array is created using an **array literal**:
*[3]int {4, 9, 6}*
+ If you store an index that is not valid for an array in a variable, and then try to access an array element using that variable as an index, you will get a panic - a runtime error.
+ You can get the number of elements in an array with the built-in *len* function.
+ You can conveniently process all the elements of an array using the special *for ..range* loop syntax, which loops through each element and assigns its index and value to variables you provide.
+ When using a *for...range* loop, you can ignore either the index or value for each element by assigning it to the _ blank identifier
+ The *os.Open* function opens a file. It returns a point to an *os.File* value representing that opened file.
+ Passing an *os.File* value to *bufio.NewScanner* returns a *bufio.Scanner* value whose *Scan* and *Text* methods can be used to read a line at a time from the file as strings.