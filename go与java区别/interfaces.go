package main

import "fmt"

// 定义不同的类型
type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

// 给猫实现接口
func (c Cat) Speak() string {
	return c.Name + "喵喵喵"
}

// 给狗实现接口
func (d Dog) Speak() string {
	return d.Name + "汪汪汪"
}

func main() {

	var s Speaker

	s = Cat{Name: "小猫"}
	fmt.Println(s.Speak())

	s = Dog{Name: "小狗"}
	fmt.Println(s.Speak())

}
