// All go files start with package declaration
// package main means this file is compiled as an executable
package main

// libary import. here we're using the built in fmt library go provides
import "fmt"

// you have to define a main function in a package main so go knows where to start the execution
// at runtime
func main() {
	fmt.Println("Hello World")
}
