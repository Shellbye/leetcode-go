package main

func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	p := 0
	// 第 0 个的处理比较特殊，因为它不需要照顾前面的数，所以直接改小自己就行
	if nums[0] > nums[1] {
		nums[0] = nums[1]
		p++
	}
	for i := 1; i < len(nums)-1; i++ {
		// 当前数大于后一个数，破坏了非递减
		if nums[i] > nums[i+1] {
			p++
			if p > 1 {
				return false
			}
			// 改动 nums[i] 变小还是改动 nums[i+1] 变大需要做一个判断
			// i 和 i-1 的关系是确定的， i >= i - 1
			// i 也大于 i + 1，现在需要确定 i + 1 和 i - 1 的关系
			// 如果 i + 1 大于 i - 1，则调小 i 和 调大 i + 1 都行
			// 但是这里还是要优先调小 i，因为 i + 1 大了更不利于非减的要求，它可能大于后面的数了
			// 如果 i + 1 小于 i - 1，则必须调大 i + 1

			if nums[i+1] >= nums[i-1] {
				nums[i] = nums[i-1]
			} else {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}

func checkPossibilityV4(nums []int) bool {
	p := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if p != -1 {
				return false
			}
			p = i
		}
	}
	return p == -1 || p == 0 || p == len(nums)-2 || nums[p-1] <= nums[p+1] || nums[p] <= nums[p+2]
}

func checkPossibilityV5(nums []int) bool {
	i := 0
	j := len(nums) - 1
	for i < j && nums[i] <= nums[i+1] {
		i++
	}
	for i < j && nums[j] >= nums[j-1] {
		j--
	}

	head := -100000
	if i != 0 {
		head = nums[i-1]
	}
	tail := 100000
	if j != len(nums)-1 {
		tail = nums[j+1]
	}

	if j-i <= 1 && (head <= nums[j] || tail >= nums[i]) {
		return true
	} else {
		return false
	}
}

func main() {
	//var nums = []int{4, 2, 3}
	//nums := []int{3, 4, 2, 3}
	nums := []int{2, 3, 3, 2, 4}
	println(checkPossibility(nums))
}
