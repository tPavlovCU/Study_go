package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	symbols := make(map[rune]int, 30)
	res := 0
	start := 0
	end := 0

	for now, symbol := range s {
		if last, ok := symbols[symbol]; ok {
			symbols[symbol] = now
			if now-last > res {
				start = last + 1
				end = now
				res = now - last
			}
		} else {
			res++
			end++
			symbols[symbol] = now
		}
	}
	fmt.Println(string(s[start:end]))
	return res

}

func main() {
	test1 := "abcabcbb"
	test2 := "bbbbb"
	test3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(test1))
	fmt.Println(lengthOfLongestSubstring(test2))
	fmt.Println(lengthOfLongestSubstring(test3))
}
