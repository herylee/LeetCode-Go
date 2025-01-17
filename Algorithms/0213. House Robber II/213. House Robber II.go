package leetcode

func rob_213(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// 由于首尾是相邻的，所以需要对比 [0，n-1]、[1，n] 这两个区间的最大值
	return max(rob_213_(nums, 0, n-2), rob_213_(nums, 1, n-1))
}

func rob_213_(nums []int, start, end int) int {
	preMax := nums[start]
	curMax := max(preMax, nums[start+1])
	for i := start + 2; i <= end; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}
