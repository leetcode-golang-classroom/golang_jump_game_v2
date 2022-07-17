package sol

func jump(nums []int) int {
	nLen := len(nums)
	left, right := 0, 0
	res := 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right < nLen-1 { // not reach right
		farthest := 0 // current farthest at current left, right
		for pos := left; pos <= right; pos++ {
			farthest = max(farthest, pos+nums[pos])
		}
		left = right + 1
		right = farthest
		res++
	}
	return res
}
