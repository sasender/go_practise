package main

import "fmt"

func main() {
	s := []int{12, 45, 7, 89, 23}
	// Find the list number
	min := s[0]
	for _, num := range s {
		if num < min {
			min = num
		}
	}
	fmt.Println("The smallest number is:", min)

}
