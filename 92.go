package main

import "fmt"

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
// Related Topics 链表
// 👍 741 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.

 */
func main() {
	node := &ListNode{}
	node.Val = 1

	node2 := &ListNode{}
	node2.Val = 2

	node3 := &ListNode{}
	node3.Val = 3

	node.Next = node2
	node2.Next = node3

	fmt.Println(reverseBetween(node, 2, 3).Next.Next.Val)

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	rightNode := prev
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := prev.Next
	curr := rightNode.Next

	prev.Next = nil
	rightNode.Next = nil

	reverseLList(leftNode)

	prev.Next = rightNode
	leftNode.Next = curr

	return dummy.Next
}

func reverseLList(head *ListNode) *ListNode {

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

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}
