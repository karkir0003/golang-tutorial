package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("error message") // creating a new error
	fmt.Println(err)
}
