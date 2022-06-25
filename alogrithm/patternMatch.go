package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) (ans []string) {
	match := func(word string) bool {
		ab, ba := map[byte]byte{}, map[byte]byte{}
		for i := 0; i < len(word); i++ {
			if v, ok := ab[word[i]]; ok { //哈希表中是否存在
				fmt.Println("ab:", string(v))
				if v != pattern[i] { //若存在哈希表中，则判断是否与当前位置相同
					return false
				}
			} else if vp, okp := ba[pattern[i]]; okp { //对比word[i]
				if vp != word[i] {
					return false
				}
			} else {
				ab[word[i]] = pattern[i] //映射，word中的第i位的字映射到pattern的i位的字母
				ba[pattern[i]] = word[i] //映射，pattern中的第i位映射到word中的第i位
			}
		}
		return true
	}
	for _, word := range words {
		if match(word) {
			ans = append(ans, word)
		}
	}
	return
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
