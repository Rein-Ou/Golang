package main

import "fmt"

func main() {
	myitem := make([]map[int]int, 1)
	itemofmap := map[int]int{1: 5}
	fmt.Println(itemofmap[1])
	myitem[0] = itemofmap
	fmt.Println(myitem[0])
	//分配切片
	items := make([]map[int]int, 5)
	for i := range items {
		//分配切片中的map元素
		//切片中i位置的元素
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A:Value of items:%v\n", items)

	//不推荐的做法
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		//出循环之后item的值丢失了
		item = make(map[int]int, 1)
		item[1] = 2
	}

	fmt.Printf("Version B: Value of items:%v\n", items2)
	fmt.Println(items2[1][1])
}
