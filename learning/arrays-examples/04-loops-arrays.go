package main

import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(arr)
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Printf("Element at index %d is %d\n", i, arr[i])
// 	}
// }
// In this code, we have declared an array arr with 5 integer elements. We then use a for loop to iterate over the array. The loop runs from 0 to the length of the array (len(arr)), and in each iteration, it prints the index and the corresponding element of the array. The output will show each element of the array along with its index.

// For loop with range
// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	for index, value := range arr {
// 		fmt.Printf("Element at index %d is %d\n", index, value)
// 	}
// }

// In this code, we use a for loop with the range keyword to iterate over the array arr. The range loop provides both the index and the value of each element in the array. We print the index and the corresponding value for each element. The output will be the same as the previous example, showing each element of the array along with its index.
// Note: The range loop is often more concise and easier to read when you need both the index and the value of the elements in the array.

//Range Loop (Values Only)
// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	for _, value := range arr {
// 		fmt.Println(value)
// 	}
// }

// In this code, we use a for loop with the range keyword to iterate over the array arr. However, we use an underscore (_) to ignore the index since we are only interested in the values of the elements. The loop will print each value in the array on a new line. The output will be:

// Range Loop (Index Only)
// func main() {
// 	arr := [5]int{20, 30, 45, 50, 60}
// 	for i := range arr {
// 		fmt.Printf("Index:%d, Value:%d\n", i, arr[i])
// 	}
// }

// In this code, we use a for loop with the range keyword to iterate over the array arr. We only capture the index (i) in the loop, and we access the value of each element using arr[i]. The loop will print both the index and the corresponding value for each element in the array. The output will show each index along with its corresponding value.

// workig-Eample-code

func main() {
	arr := [5]string{"Go", "is", "a", "great", "language"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Range Loop with Strings
	for i, value := range arr {
		fmt.Printf("%d: %s\n", i, value)
	}

	for _, value := range arr {
		fmt.Print(value + " ")
	}

}

// In this code, we have declared an array of strings arr with 5 elements. We then use a for loop to iterate over the array and print each element. Next, we use a range loop to print both the index and the value of each element in the array. Finally, we use another range loop to print only the values of the elements in a single line, separated by spaces. The output will show each string in the array along with its index, followed by all the strings printed on one line.
