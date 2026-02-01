# Go Basics

## Overview of Topics:
* [What is Go?](#what-is-go)
* [Compile and run Go program](#compile-and-run-go-program)
* [Basic hello world program](#basic-hello-world-program)
* [Data Types](#data-types)
* [Variables](#variables)
* Functions
* Control (if/else)
* Loops

## What is Go?
* Programming language invented at Google
* Characteristics
  * Static typing
  * Compiled language
  * **Concurrency support**. Golang does this through `goroutines` and channels. Goroutines are the functions/methods 
  that run concurrently with other functions or methods. Memory optimized. Channels facilitate communication between goroutines
  * Robust standard libraries for I/O, text processing graphics, networking, etc

* Overall, Provide the basics well without all the added complexity hell of classes, inheritance, bloat, etc

## Compile and Run Go program
* Produce compiled file: `go build <filename>.go` (eg: `go build helloworld.go` produces a compiled binary of `helloworld` that's executable). You can run the executable using `./helloworld`. Use this for re-running the same binary every time with no unexpected changes
* `go run <filename>.go` (eg: `go run helloworld.go`) builds the compiled binary, runs it and then removes the compiled binary.

## Basic Hello World program
> Look at `basics/helloworld.go]` for the code. We have this snippet
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```
* Package declarations are either executables (meant to be run) or libraries (other `*.go` files import them). `package main` is a special package that defines an executable program. `package main` looks for a `func main()` for running the program. 
* If I changed the package declaration in the go code above to `package hello`, then this file will be compiled as a library instead of an executable
* You have to import ONLY the packages you need otherwise you get a compilation error. This is to prevent unused packages from accumulating as the program evolves

### Good tools to have
* `gofmt` to auto format the source code to Go's standards. You can run a command like `gofmt -w helloworld.go`
* `goimports` to automanage the imports and keep only what's needed. Install [here](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) using the terminal command
* `go doc` to lookup documentation of a function in go. Eg: for the `Println` function, you can run `go doc fmt Println`

## Data Types

> Look at `basics/data-types/data-type-operations.go`

* Integers: you have `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, and `int64`. For example, `uint16` means unsigned 16 bit integer
  * The alias types (more common) are `byte` (for `uint8`), `rune` (same as `int32`). `int` is signed integer and you use this for the most part. This is machine dependent based on the architecture
* Floating point numbers: `float32` or `float64`. You can use `complex64` and `complex128` for imaginary numbers. Most of the time, use `float64`
* Arithmetic: add, subtract, multiply, divide, modulo
* Strings: Create with double quote or backticks. Double quoted strings don't contain newlines but you can use `\n` and `\t` as escape characters
* Booleans: True/False. `&&` is logical AND, `||` is logical OR, `!` is logical NOT


## Variables
> Look at `variables/variables.go` for the sample code

* Variable is just a memory address for storing a value. We specify type and name

Example Declaration:
```go
var x string
x = "Hello, World"
fmt.Println(x)
```

As you may guess `x` is name of variable. It's of type string and we assign it the value of `"Hello, World"`

The preferred way is through this shorthand
```go
x := "Hello, World"
fmt.Println(x)
```
Here, we see that variable `x` is assigned value of `"Hello, World"`. Go compiler autoinfers the type

Use camel case for naming variables in go

**NOTE:** "Go is lexically scoped using blocks". That means if you declare a variable in a function, the variable's scope is that function. Here's an example

```go
var x string = "Hello, World"

func main() {
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```
The `x` variable here can be used in both `main()` and `f()` functions

```go
func main() {
    var x string = "Hello, World"
    fmt.Println(x)
}

func f() {
    fmt.Println(x)
}
```
However, we will get a compilation error in this code when we move `x` into the scope of the `main()` function. Function `f()` can't access `x`

You can also define constants which are variables where we cannot change the value

Use `fmt` library also for parsing input from the terminal. See `variables/double_number.go` for a sample script

