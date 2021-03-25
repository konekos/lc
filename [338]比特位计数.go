package main
//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
//
// 示例 1:
//
// 输入: 2
//输出: [0,1,1]
//
// 示例 2:
//
// 输入: 5
//输出: [0,1,1,2,1,2]
//
// 进阶:
//
//
// 给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
// 要求算法的空间复杂度为O(n)。
// 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
//
// Related Topics 位运算 动态规划
// 👍 546 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0,1}
	}

	arr := make([]int,0)
	arr = append(arr, 0)
	arr = append(arr, 1)
	arr = append(arr, 1)

	m := make(map[int]int, num)
	m[1] = 1
	m[2] = 1

	for i := 3; i <= num; i++ {
		if i % 2==0 {
			m[i] = m[i/2]
			arr = append(arr, m[i])
		}else {
			m[i] = m[i-1] + 1
			arr = append(arr, m[i])
		}
	}

	return arr

}

func main() {

}
//leetcode submit region end(Prohibit modification and deletion)
