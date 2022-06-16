package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res, short := "", ""
	if len(strs) > 0 {
		res, short = strs[0], strs[0]
	}
	for _, str := range strs {
		if len(str) < len(short) {
			short, str = str, short
		}
		temp := short
		for v := range short {
			if str[v] != short[v] {
				temp = short[:v]
				break
			}
		}
		if len(temp) < len(res) {
			res = temp
		}
	}
	return res
}

func main() {
	teststring := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(teststring))
}
