package main

import (
	"fmt"
)

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
}

func deferFunc() int {
	fmt.Println("defer func")
	return 0
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()

}

func main() {

	returnAndDefer()

}
