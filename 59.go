package main //给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
// Related Topics 数组
// 👍 332 👎 0

func main() {

	generateMatrix(3)

}

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	tv := make([][]int, 0)
	for i := 0; i < n; i++ {
		arr := make([]int, n)
		tv = append(tv, arr)
	}

	num := 1
	// 偶数时直接填充完毕，奇数时中间值是 n2
	for i := 0; i < n/2; i++ {
		// top
		for t := i; t < n-i; t++ {
			tv[i][t] = num
			num++
		}
		// right
		for r := i + 1; r < n-i; r++ {
			tv[r][n-i-1] = num
			num++
		}
		// bottom
		for b := n - i - 2; b >= i; b-- {
			tv[n-i-1][b] = num
			num++
		}
		// left
		for l := n - i - 2; l > i; l-- {
			tv[l][i] = num
			num++
		}
	}
	// 奇数时填充正中间
	if n%2 == 1 {
		tv[n/2][n/2] = n * n
	}
	return tv
}

//leetcode submit region end(Prohibit modification and deletion)
