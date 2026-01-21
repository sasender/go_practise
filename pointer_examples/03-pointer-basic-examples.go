// Program to assign memory address to pointer

package main

import "fmt"

func main() {

	var name = "Sasender"
	var ptr *string

	// assign the memory address of name to the pointer
	ptr = &name
	ptr = &name

	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Address of the variable", &name)
	fmt.Println("Value of the variable", *ptr)

}
