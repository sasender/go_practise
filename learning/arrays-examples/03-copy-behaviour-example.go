package main

import "fmt"

var arr1 = [5]int{1, 2, 3, 4, 8}
var arr2 = arr1

func main() {
	arr2[4] = 5
	fmt.Println(arr1)
	fmt.Println(arr2)
}

// In this code, we have declared an array arr1 with 5 integer elements. We then create another array arr2 and assign it the value of arr1. This means that arr2 is a copy of arr1, and they are two separate arrays in memory. When we modify arr2 by changing the value at index 4 to 5, it does not affect arr1. Therefore, when we print both arrays, we will see that arr1 remains unchanged as [1 2 3 4 8], while arr2 reflects the change as [1 2 3 4 5].
