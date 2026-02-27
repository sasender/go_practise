package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func userage(p *Person, newAge int) {
	p.Age = newAge
	fmt.Println("Value of p:", &p)
}

func main() {
	user := Person{Name: "Sasender", Age: 25}
	fmt.Println("Before:", user.Name, user.Age)
	userage(&user, 30)
	//fmt.Println("Value of user:", &user)
	fmt.Println("After:", user.Name, user.Age)
}
