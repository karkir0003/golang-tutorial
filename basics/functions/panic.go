package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str) // print the panic error message!
	}()
	panic("PANIC") // no error raised but we run the defer even in panic
}
