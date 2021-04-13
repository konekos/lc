package main

import "math"

//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bs
//t/ 相同
//
//
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,6,1,3]
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [2, 100] 内
// 0 <= Node.val <= 105
//
//
//
// Related Topics 树 深度优先搜索 递归
// 👍 154 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt32
	arr := make([]int, 0)
	var mid func(root *TreeNode)

	mid = func(root *TreeNode) {
		if root == nil {
			return
		}
		mid(root.Left)
		arr = append(arr, root.Val)
		mid(root.Right)
	}
	mid(root)

	for i := 0; i < len(arr)-1; i++ {
		res = min(res, int(math.Abs(float64(arr[i+1]-arr[i]))))
	}
	return res
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
