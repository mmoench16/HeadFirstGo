# Head First Go - Notes - Chapter 1

## Note 1

### Function Calls

A function is a chunk of code that you can call from other places in your program.

When calling a function, you can use arguments to provide the function with data.

## Note 2

### Types

Values in Go are classified into different types. which specify what the values can be used for. 

Math operations and comparisons between different types are not allowed, but you can convert a value to a new type if needed.

Go variables can only store values of their declared type.

## Bullet Points

+ A **package** is a group of related functions and other code.
+ Before you can use a package's functions within a Go file, you need to **import** that package.
+ A *string* is a series of bytes that usually represent text characters.
+ A *rune* represents a single text character.
+ Go's two most common numeric types are *int*, which holds integers, and *float64*, which holds floating-point numbers.
+ The *bool* type holds Boolean values, which are either *true* or *false*
+ A **variable** is a piece of storage that can contain values of a specified type.
+ If no value has been assigned to a variable, it will contain the **zero value** for its type. Examples of zero values include *0* for *int* or *float64* variables, or *""* for *string* variables
+ You can declare a variable and assign it a value at the same time using a *:=* **short variable declaration**.
+ A variable, function, or type can only be accessed from code in other packages if its name begins with a capital letter.
+ The *go fmt* command automatically reformats source files to use Go standard formatting. You should run *go fmt* on any code that you plan to share with others.
+ The *go build* command **compiles** Go source code into a binary format that computers can execute.
+ the *go run* command compiles and runs a program without saving an executable file in the current directory.