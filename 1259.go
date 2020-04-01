package main

import "fmt"

func findNumbers(nums []int) int {
	var ret = 0
	for _, value := range nums {
		if value > 10 && value < 100 {
			ret++
		}
		if value > 1000 && value < 10000 {
			ret++
		}
	}
	return ret
}

func main() {
	var nums = []int{3, 2, 0, 4}
	fmt.Println(findNumbers(nums))
}
