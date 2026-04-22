package main

import "fmt"

func main() {
	cpuUsage := 85
	if cpuUsage > 80 {
		fmt.Println("High usage")
	} else if cpuUsage >= 50 && cpuUsage <= 80 {
		fmt.Println("Moderate")
	} else if cpuUsage < 50 {
		fmt.Println("Low")
	}
}
