package main

import (
	"fmt"
	"time"
)

//defer 会在函数返回前执行 栈 先进后出

func deferFunc() int {
	fmt.Println("defer func0")
	return 0
}

func deferFunc1() int {
	fmt.Println("defer func1")
	return 0
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()
	defer deferFunc1()
	defer func() int {
		fmt.Println("defer func2")
		return 10
	}()
	time.Sleep(1 * time.Second)
	return returnFunc()

}

func main() {

	returnAndDefer()

}
