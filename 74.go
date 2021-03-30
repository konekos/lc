package main

import (
	"fmt"
	"sort"
)

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104
//
// Related Topics 数组 二分查找
// 👍 360 👎 0

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{10, 11, 16, 20}
	arr3 := []int{23, 30, 34, 50}
	matrix := [][]int{arr1, arr2, arr3}
	sort.Search(4, func(i int) bool {

	})
	println(searchMatrix(matrix, 30))
}

func searchMatrix(matrix [][]int, target int) bool {

	n := len(matrix)

	arr1 := make([]int, 0)
	for i := 0; i < n; i++ {
		arr1 = append(arr1, matrix[i][0])
	}

	y := binary(arr1, target)
	if y < 0 {
		return false
	}

	fmt.Println(y)
	return binary2(matrix[y], target)
}

func binary(arr []int, target int) (index int) {
	x := 0
	y := len(arr) - 1

	if len(arr) == 1 {
		if arr[0] <= target {
			return 0
		}
		return -1
	}

	for {

		mid := (x + y) / 2
		if arr[mid] == target {
			return mid
		}

		if arr[x] == target {
			return x
		}
		if arr[y] <= target {
			return y
		}

		if y-x == 1 {
			if arr[x] < target && arr[y] > target {
				return mid
			}
			return -1
		}

		if arr[mid] > target {
			y = mid
		}
		if arr[mid] < target {
			x = mid
		}

	}
}

func binary2(arr []int, target int) bool {
	x := 0
	y := len(arr) - 1

	for {

		mid := (x + y) / 2

		if arr[x] == target || arr[y] == target || arr[mid] == target {
			return true
		}

		if y-x == 1 || y == x {
			return false
		}

		if arr[mid] > target {
			y = mid
		}
		if arr[mid] < target {
			x = mid
		}

	}
}
