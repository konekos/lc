package main //给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
//
// 我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
//
//
//
// 示例 1:
//
// 输入: nums = [4,2,3]
//输出: true
//解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
//
//
// 示例 2:
//
// 输入: nums = [4,2,1]
//输出: false
//解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
//
//
//
//
// 说明：
//
//
// 1 <= n <= 10 ^ 4
// - 10 ^ 5 <= nums[i] <= 10 ^ 5
//
// Related Topics 数组
// 👍 473 👎 0

func main() {
	var f bool
	println(f)
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkPossibility(nums []int) bool {
	var p int
	if len(nums) < 3 {
		return true
	}
	var f bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			p = i
			f = true
		}
	}

	if !f {
		return true
	}

	var tmp = nums[p]
	var tmp1 = nums[p+1]

	var b1 bool = true
	var b2 bool = true

	nums[p+1] = tmp

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			b1 = false
		}
	}

	nums[p] = tmp1
	nums[p+1] = tmp1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			b2 = false
		}
	}

	return b1 || b2

}

//leetcode submit region end(Prohibit modification and deletion)
