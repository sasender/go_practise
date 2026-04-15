package main

import "fmt"

func main() {
	var age int = 30
	var name string = "Sasender"
	var isStudent bool = false
	var height float64 = 5.9

	fmt.Println(age, name, isStudent, height)
	fmt.Printf("Name: %s, Age: %d, Is Student: %t, Height: %.1f\n", name, age, isStudent, height)

}
