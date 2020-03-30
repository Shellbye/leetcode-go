package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if nums[j] < nums[i] {
				ret[i] += 1
			}
		}
	}
	return ret
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
