package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 回溯算法
// 👍 650 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	track := make([]int, 0)
	vis := make([]bool, n)

	var backtrack func(idx int)

	backtrack = func(idx int) {
		if idx == n {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}

		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			track = append(track, nums[i])
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

	arr := permuteUnique([]int{1, 2, 2})
	for _, ints := range arr {
		fmt.Printf("%v", ints)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
