package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "This is a string"
	string2 := "It's a string"

	fmt.Printf("%t\n", strings.HasSuffix(string1, "string"))
	fmt.Printf("%t\n", strings.HasSuffix(string2, "string"))
	fmt.Printf("%t\n", strings.HasPrefix(string1, "Th"))
	fmt.Printf("%t\n", strings.HasPrefix(string2, "Th"))
	string3 := "nishuo什么东西nishuo"
	//Index,LastIndex,子字符串或字符在父字符串中出现的位置
	fmt.Println(strings.IndexRune(string3, rune('什')))
	//只替换第一次出现的
	fmt.Println(strings.Replace(string3, "nishuo", "你说", 1))
	fmt.Println(strings.Count(string3, "nishuo"))
	fmt.Println(strings.Repeat(string2, 2))
	fmt.Println(strings.ToLower(string1))
	fmt.Println(strings.ToUpper(string2))
	strings.Trim(string3, "nishuo")
	strings.TrimSpace(string1)
	strings.TrimLeft(string3, "nishuo")
	//按空格分割,分割出来返回的是slice
	sl := strings.Fields(string1)
	for _, val := range sl {
		fmt.Println(val)
	}
	//自定义分隔符
	sl2 := strings.Split(string2, "'")
	for _, val := range sl2 {
		fmt.Println(val)
	}

	string4 := "你在说什么"
	reader := strings.NewReader(string4)
	values, _ := reader.ReadByte()
	character, size, _ := reader.ReadRune()

	fmt.Println(values)
	fmt.Printf("%c - %d", character, size)
}
