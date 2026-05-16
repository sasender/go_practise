package main

import "fmt"

func main() {
	pods := []string{
		"Running",
		"Pending",
		"Failed",
	}
	for _, status := range pods {
		fmt.Println(status)
	}
}
