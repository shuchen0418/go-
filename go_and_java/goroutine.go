package main

import (
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

/* func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	x := <-ch
	// y := <-ch
	// z := <-ch

	fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	//互斥锁
	mu.Lock()
	//其他goroutine在这里阻塞，直到锁释放
	mu.Unlock()

	//WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 1")
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine 2")
	}()

	wg.Wait()

} */
