package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	lastNonZero := 0
	zeroes := 0
	for i := 0; i < len(nums); i++ {
		el := nums[i]
		if el != 0 {
			nums[lastNonZero] = el
			lastNonZero++
		} else {
			zeroes++
		}
	}

	for m := range zeroes {
		nums[len(nums)-m-1] = 0
	}
	fmt.Println(nums)

}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
