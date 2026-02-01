# Go Basics

## Overview of Topics:
* [What is Go?](#what-is-go)
* [Compile and run Go program](#compile-and-run-go-program)
* [Basic hello world program](#basic-hello-world-program)
* Variables
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
* Look at `basics/helloworld.go` for the code. We have this snippet
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
* `gofmt` to auto format the source code to Go's standards 
* `goimports` to automanage the imports and keep only what's needed

