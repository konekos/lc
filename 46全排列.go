package main

import "fmt"

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
// 示例:
//
// 输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// Related Topics 回溯算法
// 👍 1254 👎 0

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}

func add(arr []int) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {

	res := make([][]int, 0)
	track := make([]int, 0)

	var backtrack func()

	backtrack = func() {
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}
	outer:
		for i := 0; i < len(nums); i++ {
			for _, num := range track {
				if num == nums[i] {
					continue outer
				}
			}
			track = append(track, nums[i])
			backtrack()
			track = track[:len(track)-1]
		}
	}

	backtrack()
	return res
}

//func backtrack(nums []int, track []int) {
//	if len(track) == len(nums) {
//		res = append(res, track)
//		return
//	}
//	outer:
//	for i := 0; i < len(nums); i++ {
//		for _, num := range track {
//			if nums[i] == num {
//				continue outer
//			}
//		}
//
//		track = append(track, nums[i])
//
//		backtrack(nums, track)
//
//		track = track[:len(track) - 1]
//	}
//
//	return
//}

//leetcode submit region end(Prohibit modification and deletion)
