package main

import "fmt"

func main() {
	cpu := []int{30, 85, 60, 95, 20}
	for _, cpu := range cpu {
		if cpu > 80 {
			fmt.Println(cpu, "high cpu usage alert ! >80%")
		}
	}
}
