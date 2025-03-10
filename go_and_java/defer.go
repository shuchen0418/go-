package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

//前两个defer是在注册时就固定了a的值(值传递)，
// 而第三个defer是在执行时才查找a的当前值(引用传递)

func test(a int) { //无返回值函数
	defer fmt.Println("1、a =", a)                    //方法
	defer func(v int) { fmt.Println("2、a =", v) }(a) //有参函数
	defer func() { fmt.Println("3、a =", a) }()       //无参函数
	a++
}

// 无名返回值：defer不能修改返回值（因为返回的是值的副本）
func test1() int { //无名返回值函数
	var a int
	defer func() {
		a++
		fmt.Println("defer1:", a)
	}()
	defer func() {
		a++
		fmt.Println("defer2:", a)
	}()
	return a
}

// 有名返回值：defer可以修改返回值（因为操作的是同一个命名变量）
func test2() (a int) { //有名返回值函数
	defer func() {
		a++
		fmt.Println("defer1:", a)
	}()

	defer func() {
		a++
		fmt.Println("defer2:", a)
	}()
	return a
}

func test3() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func f() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Println("r =", r)
	}(r)
	fmt.Println("r =", r)
	return 1
}

type Test struct {
	name string
}

// 方法
func (t *Test) pp() {
	fmt.Println(t.name)
}

func pp(t Test) {
	fmt.Println(t.name)
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	//多个goroutine同时读写map，需要加锁
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

/* func main() {
	// test(1)
	// fmt.Println("return :", test1())
	// fmt.Println("return :", test2())
	// fmt.Println("return :", test3())
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer pp(t)
	}
} */
