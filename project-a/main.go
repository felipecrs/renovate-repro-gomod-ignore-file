package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Hello, World!")
	err := errors.New("an error")
	fmt.Println(err)
}
