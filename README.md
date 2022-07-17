# golang_jump_game_v2

Given an array of non-negative integers `nums`, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.

## Examples

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

```

**Example 2:**

```
Input: nums = [2,3,0,1,4]
Output: 2

```

**Constraints:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 1000`

## 解析

給定一個非負整數陣列 nums

每一個元素 nums[i] 代表該 index 所能夠 jump 的最大值

必須要從 index 0 開始 jump 到最後一個 index

要求寫一個演算法計算從 0 到最後一個 index 所需最小的 jump 個數

假設透過 Dynamic Programming 

定義 dp[i] = 到 index 所需要的最少 jump 數

可以發現 

dp[pos+jump] = min(1 + dp[pos], dp[pos+jump])  for all  1 ≤ jump ≤ nums[pos]  if dp[pos+jump] ≠ 0

dp[pos+jump] = 1+dp[pos] if dp[pos+jump] == 0 for all  1 ≤ jump ≤ nums[pos] 

透過 loop 過所有 nums 的 index 還有 每個 max = nums[pos]

所求是 dp[nLen-1]

時間複雜度是 O($n^2$)

空間複雜度是 O(n)

透過 Greedy Algorithm

可以發現每次 jump

都會有一個 range 如下圖

![](https://i.imgur.com/PscMqch.png)

每次透過 nums[pos] 都可以找到每次在 pos 可以跳的的最遠距離

如同在做BFS 

當最遠可以跳到最後一個 index 則完成

所以每次只要紀錄當下 pos 能跳的最遠距離當作 right , 而上一次的最遠距離+1 的做 left

當 right 到達 len(nums) - 1 就結束了

每次針對 index 從 left 到 right

計算 farthest = max(farthest, pos + nums[pos])

然後更新 left = right+1

更新 right = farthest

然後每更新一次 left, right 就累加 res +1

當 right 到達 len(nums) - 1 時 , res 及為所求

可以看到每次都是透過 逼近 right 來做處理

時間複雜度是 O(n)

空間複雜度是 O(1) 

## 程式碼
```go
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

```
## 困難點

1. 要理解到如何透過 pos 的遞近關係來做推演出 greedy 的作法

## Solve Point

- [x]  初始化 left = 0, right = 0, 因為一開始只有起始點 0
- [x]  初始化 farthest = 0, 透過 farthest = max(farthest, pos+ nums[pos]) 來推算 在可能的 pos 之下的 farthest
- [x]  更新 left = right + 1, right =farthest , res += 1 來推算下一個 邊界
- [x]  當 right = len(nums) -1 回傳 res