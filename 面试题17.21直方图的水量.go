package main

import "fmt"

//给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 感谢 Marco
//s 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
// Related Topics 栈 数组 双指针
// 👍 97 👎 0

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	i := trap(arr)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	res := 0
	var calcute func(arr []int)
	calcute = func(arr []int) {
		n := len(arr)
		if n < 3 {
			return
		}
		i := findMax(arr)
		if i > 1 {
			left := arr[:i]
			leftM := findMax(left)
			q := left[leftM+1 : i]
			res += calres(q, left[leftM])
			if leftM > 1 {
				calcute(arr[:leftM+1])
			}
		}

		if i < n-2 {
			right := arr[i+1:]
			index := findMax(right)
			rightM := i + index + 1
			q := right[:index]
			res += calres(q, right[index])

			if rightM < n-2 {
				calcute(arr[rightM:])
			}
		}

	}
	calcute(height)
	return res

}

func findMax(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}

func calres(arr []int, max int) int {

	res := 0
	for i := 0; i < len(arr); i++ {
		res += max - arr[i]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
