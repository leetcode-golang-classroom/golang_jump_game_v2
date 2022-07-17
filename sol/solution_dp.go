package sol

func jump_dp(nums []int) int {
	nLen := len(nums)
	dp := make([]int, nLen)
	dp[0] = 0
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for pos := 0; pos < nLen-1; pos++ {
		maxJump := nums[pos]
		for jp := maxJump; jp >= 1; jp-- {
			if jp+pos <= nLen-1 {
				if dp[jp+pos] == 0 {
					dp[jp+pos] = 1 + dp[pos]
				}
				dp[jp+pos] = min(dp[jp+pos], pos+nums[pos])
			}
		}
	}
	return dp[nLen-1]
}
