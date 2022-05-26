//当一个变量被声明之后，系统自动赋予它该类型的零值
//若你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写
//:=只能被用于函数体内，不可以用于全局变量的声明与赋值
package main

import (
	"fmt"
	"math"
	"os"
)

var Pi float64

func main() {
	PATH := os.Getenv("PATH")
	fmt.Println(PATH)
	fmt.Println(Pi)
}

//优先级比main高,一个源文件只能包含一个init函数
func init() {
	Pi = 4 * math.Atan(1)

}
