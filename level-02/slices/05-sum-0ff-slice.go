package main

import "fmt"

// func main() {
// 	s := []int{10, 20, 30}
// 	sum := 0
// 	for _, value := range s {
// 		sum += value
// 	}
// 	fmt.Println(sum)
// }

// Check Number Exists or Not

// func main() {
// 	s := []int{10, 20, 30, 40}
// 	target := 30
// 	found := false
// 	for _, value := range s {
// 		if value == target {
// 			found = true
// 			break
// 		}
// 	}
// 	if found {
// 		fmt.Println("Number found in the slice. found =", found)
// 	} else {
// 		fmt.Println("Number not found in the slice, found =", found)
// 	}

// }

// ## Remove Duplicates from Slice

func main() {
	s := []int{1, 2, 2, 3, 4, 4, 5}
	unique := make(map[int]bool)
	for _, value := range s {
		unique[value] = true
	}
	result := []int{}
	for value := range unique {
		result = append(result, value)
	}
	fmt.Println(result)
}
