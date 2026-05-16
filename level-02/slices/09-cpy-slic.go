package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}

	s2 := make([]int, len(s1))
	copy(s2, s1)

	fmt.Println(s1)
	fmt.Println(s2)
}
