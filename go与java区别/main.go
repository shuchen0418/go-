package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Speaker interface {
	Speak() string
}

func (p Person) Greet() {
	fmt.Println("Hello my name is ", p.Name)
}

func main() {
	p := Person{Name: "奥特曼", Age: 19}
	p.Greet()

	// 使用 Cat 和 Dog
	cat := Cat{Name: "小猫"}
	dog := Dog{Name: "小狗"}

	var s Speaker
	s = cat
	fmt.Println(s.Speak())

	s = dog
	fmt.Println(s.Speak())
}
