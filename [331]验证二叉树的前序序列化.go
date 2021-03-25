package main

import (
	"fmt"
	"strings"
)

//序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
//
//      _9_
//    /   \
//   3     2
//  / \   / \
// 4   1  #  6
/// \ / \   / \
//# # # #   # #
//
//
// 例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
//
// 给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
//
// 每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
//
// 你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
//
// 示例 1:
//
// 输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
//输出: true
//
// 示例 2:
//
// 输入: "1,#"
//输出: false
//
//
// 示例 3:
//
// 输入: "9,#,#,1"
//输出: false
// Related Topics 栈
// 👍 234 👎 0

func main() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)

	fmt.Println(arr[:3])
	fmt.Println(arr[:9])

	arr = arr[:9]

	fmt.Println(arr)

	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isValidSerialization(preorder string) bool {
	arr := strings.Split(preorder, ",")
	if arr[0] == "#" && len(arr) == 1 {
		return true
	}

	if arr[0] == "#" && len(arr) > 1 {
		return false
	}


	ind := 0
	outd := 2

	for i := 1; i < len(arr); i++ {
		node := arr[i]
		if node == "#" {
			ind ++
		}else {
			outd +=2
			ind ++
		}

		if i != len(arr)-1 && ind >= outd {
			return false
		}
	}

	return ind == outd

}
