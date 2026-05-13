//find the largers-number in the slices
// we need tp print the largest numbers in above slices list

package main

import "fmt"

func main() {
	s := []int{12, 45, 7, 89, 23}

	// Find the largest number
	max := s[0]
	for _, num := range s {
		if num > max {
			max = num
		}
	}

	fmt.Println("The largest number is:", max)
}
