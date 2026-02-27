package main

import "fmt"

type Person struct {
	name string
	age  int
}

func birthday(p *Person) {
	p.age += 15
}

func main() {
	person := Person{"Alice", 20}
	birthday(&person)
	fmt.Println(person) // {Alice 35}
}
