package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello my name is ", p.Name)
}

func main() {
	p := Person{Name: "奥特曼", Age: 19}
	p.Greet()

}
