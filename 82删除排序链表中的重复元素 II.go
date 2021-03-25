package main

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//
// 返回同样按升序排列的结果链表。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序排列
//
// Related Topics 链表
// 👍 522 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -101}
	prev := dummy
	curr := head

	for curr != nil {

		next := curr.Next
		if next == nil {
			prev.Next = curr
			break
		}
		if next.Val != curr.Val {
			prev.Next = curr
			prev = curr
			curr = next
		} else {
			val := curr.Val
			prev.Next = nil
			for curr != nil && curr.Val == val {
				if curr.Next != nil {
					curr = curr.Next
				} else {
					curr = nil
				}

			}
		}

	}
	return dummy.Next

}
