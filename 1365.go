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

func smallerNumbersThanCurrentV2(nums []int) []int {
	var buckets [101]int
	for i := 0; i < len(nums); i++ {
		buckets[nums[i]] += 1
	}

	var sums [101]int
	sums[0] = 0
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + buckets[i-1]
	}

	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = sums[nums[i]]
	}

	return ret
}

func smallerNumbersThanCurrentV3(nums []int) []int {
	var buckets [101]int
	for i := 0; i < len(nums); i++ {
		buckets[nums[i]] += 1
	}

	pre := 0
	for i := 0; i < len(buckets); i++ {
		buckets[i], pre = pre, buckets[i]+pre
	}

	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = buckets[nums[i]]
	}

	return ret
}

func main() {
	nums := []int{3, 2, 0, 4}
	fmt.Println(smallerNumbersThanCurrentV3(nums))
}
