package main

import "go/ast"

//反转一个单链表。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//
// 进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// Related Topics 链表
// 👍 1588 👎 0

func main() {
	n1 := ListNode{}

	n1.Val = 1

	n2 := ListNode{}
	n2.Val = 2
	n1.Next = &n2

	listNode := reverseList(&n1)
	println(listNode.Val)
	println(listNode.Next.Val)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	phead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return phead
}

//leetcode submit region end(Prohibit modification and deletion)
