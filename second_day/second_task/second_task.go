package main

import "fmt"
import "sort"

func twoSum(nums []int, target int) []int {
	indexes := make([]int, 0, len(nums))
	for i := range nums {
		indexes = append(indexes, i)
	}
	sort.Slice(indexes, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	i := 0
	j := len(nums) - 1

	for i < j {
		left_el := nums[indexes[i]]
		right_el := nums[indexes[j]]
		sum := left_el + right_el

		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{indexes[i], indexes[j]}
		}
	}
	return nil
}

func main() {
	test := []int{3, 2, 4}
	fmt.Println(twoSum(test, 6))
}
