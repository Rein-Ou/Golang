package main

import "fmt"

func main() {
	var a = [4]float32{1, 2, 3, 4}
	fmt.Println(Sum(a))
	var b = []float32{1, 2, 3, 4}
	fmt.Println(Sum_Slice(b))
}

func Sum(arrF [4]float32) (result float32) {
	result = 0
	for _, val := range arrF {
		result += val
	}
	return
}

func Sum_Slice(arrF []float32) (result float32) {
	result = 0
	for _, val := range arrF {
		result += val
	}
	return
}

func SumAndAverage(arr []int) (int, float32) {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	average := sum / len(arr)
	return sum, float32(average)
}
