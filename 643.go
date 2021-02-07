package main

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
//
//
//
// 示例：
//
//
//输入：[1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 30,000。
// 所给数据范围 [-10,000，10,000]。
//
// Related Topics 数组
// 👍 148 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	var m int
	for i := 0; i < k; i++ {
		m += nums[i]
	}

	sub := m - nums[0]

	left := 1

	for left <= len(nums)-k {
		var tmp int
		tmp = sub + nums[left+k-1]
		sub = tmp - nums[left]
		m = max(m, tmp)
		left++
	}
	return float64(m) / float64(k)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
