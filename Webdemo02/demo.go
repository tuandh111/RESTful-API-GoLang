package main

import "fmt"

func main() {
	a := 10

	var pointer *int

	pointer = &a

	fmt.Println(*pointer)
}
