package main

import "fmt"

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109
//
// Related Topics 链表 双指针
// 👍 531 👎 0

func main() {
	fmt.Println(7 % 5)
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

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	curr := head
	length := 1
	for curr.Next != nil {
		curr = curr.Next
		length++
	}

	tail := curr

	offSet := length - k%length

	if offSet == 0 {
		return head
	}

	curr.Next = head

	for offSet > 0 {
		curr = curr.Next
		offSet--
	}

	newTail := curr
	newHead := curr.Next

	tail.Next = head

	newTail.Next = nil

	return newHead

}

//leetcode submit region end(Prohibit modification and deletion)
