package main

import "fmt"

func secondLargest(num []int) (int, bool) {
	if len(num) < 2 {

		return 0, false
	}

	first, second := num[0], num[0]
	foundDiff := false

	for _, n := range num {
		if n > first {
			second = first
			first = n
			foundDiff = true
		} else if n > second && n != first {
			second = n
			foundDiff = true
		}
	}
	return second, foundDiff
}

func main() {
	s := []int{10, 50, 20, 40}
	if second, ok := secondLargest(s); ok {
		fmt.Println("Second largest number:", second)
	} else {
		fmt.Println("Not enough distinct elements to find the second largest.")
	}
}
