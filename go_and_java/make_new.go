package main

import "fmt"

func a() {

	// - %T : 打印值的类型
	// - %d : 打印整数
	// - %s : 打印字符串
	// - %f : 打印浮点数
	// - %p : 打印指针的地址
	// - %+v : 打印结构体时会带字段名
	// - %#v : 打印值的 Go 语法表示

	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	fmt.Printf("%v\n", nums)

	nums = append(nums, 5, 6, 7)
	fmt.Printf("%v\n", nums)

	moreNums := []int{8, 9, 10}
	nums = append(nums, moreNums...)
	fmt.Printf("%v\n", nums)

	var empty []string
	empty = append(empty, "hello", "world")
	empty = append(empty, "nihao")
	fmt.Printf("%v\n", empty)

	nums = append(nums[:2], nums[3:]...)
	fmt.Printf("%v\n", nums)

	nums = append(nums[:3], append([]int{100}, nums[3:]...)...)
	fmt.Printf("%v\n", nums)

	newSlice := append([]int{}, nums...)
	fmt.Printf("%v\n", newSlice)

	a := [3]int{}
	b := new([3]int)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", *b)

}
