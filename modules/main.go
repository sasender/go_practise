package main

import (
	"fmt"
	"modules/packages"
)

func main() {
	a := 10
	b := 8
	fmt.Println("Hellow World")
	packages.PrintHello()
	packages.Cricket()
	fmt.Println(packages.Addition(a, b))
	fmt.Println(packages.Substraction(a, b))

}
