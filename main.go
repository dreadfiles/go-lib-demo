package main

import (
	"fmt"

	"github.com/dreadfiles/go-lib"
)

func main() {
	fmt.Println(lib.Add(3, 5))
	fmt.Println(lib.Substract(6, 2))

	fmt.Println(lib.Add(5, 7))
	fmt.Println(lib.Substract(7, 1))
}
