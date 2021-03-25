package main //给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
import "fmt"

//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
//
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// Related Topics 数组
// 👍 649 👎 0

// 1 2 3
// 4 5 6
// 7 8 9

func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{5, 6, 7, 8}
	a3 := []int{9, 10, 11, 12}

	arr := [][]int{a1, a2, a3}
	ints := spiralOrder(arr)
	fmt.Println(ints)

}

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	m := len(matrix[0])
	n := len(matrix)
	times := min((m+1)/2, (n+1)/2)
	fmt.Println(times)
	arr := make([]int, 0)

	for i := 0; i < times; i++ {
		// [i,i]-[i, m-i-1]
		for t := i; t < m-i; t++ {
			arr = append(arr, matrix[i][t])
		}

		// [i+1, m-i-1] - [n-i-1,m-i-1]
		f1 := false
		for r := i + 1; r < n-i; r++ {
			f1 = true
			arr = append(arr, matrix[r][m-i-1])
		}

		//[n-i-1, m-i-2]-[n-i-1, i]
		f2 := false
		for b := m - i - 2; b >= i; b-- {
			if !f1 {
				break
			}
			f2 = true
			arr = append(arr, matrix[n-i-1][b])
		}

		//[n-i-2,i]-[i,i-1]

		for l := n - i - 2; l > i; l-- {
			if !f2 {
				break
			}
			arr = append(arr, matrix[l][i])
		}
	}

	return arr
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
