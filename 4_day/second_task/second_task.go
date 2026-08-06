package main

import "fmt"

func isGood(symb uint8) bool {
	if symb <= 'z' && symb >= 'A' {
		return true
	} else {
		return false
	}
}

func equal(symb1, symb2 uint8) bool {
	res := int(symb1) - int(symb2)
	if res == 32 || res == -32 || res == 0 {
		return true
	} else {
		return false
	}
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {

		if isGood(s[i]) && isGood(s[j]) {
			if !equal(s[i], s[j]) {
				return false
			}
			i++
			j--
		} else {
			if !isGood(s[i]) {
				i++
			}
			if !isGood(s[j]) {
				j--
			}
		}
	}
	return true

}

func main() {
	test1 := "A man a plan a canal Panama"
	fmt.Println(isPalindrome(test1))
}
