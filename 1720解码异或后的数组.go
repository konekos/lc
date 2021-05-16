package main

import "fmt"

//未知 整数数组 arr 由 n 个非负整数组成。
//
// 经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，a
//rr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。
//
// 给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。
//
// 请解码返回原数组 arr 。可以证明答案存在并且是唯一的。
//
//
//
// 示例 1：
//
//
//输入：encoded = [1,2,3], first = 1
//输出：[1,0,2,1]
//解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [
//1,2,3]
//
//
// 示例 2：
//
//
//输入：encoded = [6,2,7,3], first = 4
//输出：[4,2,0,7,4]
//
//
//
//
// 提示：
//
//
// 2 <= n <= 104
// encoded.length == n - 1
// 0 <= encoded[i] <= 105
// 0 <= first <= 105
//
// Related Topics 位运算
// 👍 50 👎 0

func main() {
	fmt.Printf("%v", decode([]int{6, 2, 7, 3}, 4))
}

//leetcode submit region begin(Prohibit modification and deletion)
func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i < len(encoded)+1; i++ {
		num := encoded[i-1] ^ res[i-1]
		res[i] = num
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//leetcode submit region end(Prohibit modification and deletion)
