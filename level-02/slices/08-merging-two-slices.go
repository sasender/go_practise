package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	merged := append(s1, s2...)
	fmt.Println(merged)

}
