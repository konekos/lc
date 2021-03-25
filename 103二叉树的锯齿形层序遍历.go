package main

func main() {
	zigzagLevelOrder(&TreeNode{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var arr [][]int
	nodes := []*TreeNode{root}
	var tmpa []int
	level := 1
	for len(nodes) > 0 {
		tmpa = []int{}
		var tmp []*TreeNode
		for i := 0; i < len(nodes); i++ {
			tmpa = append(tmpa, nodes[i].Val)
			if nodes[i].Left != nil {
				tmp = append(tmp, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				tmp = append(tmp, nodes[i].Right)
			}
		}

		if level%2 == 0 {
			if len(tmpa) > 1 {
				for i := 0; i < len(tmpa)/2; i++ {
					tmp := tmpa[i]
					tmpa[i] = tmpa[len(tmpa)-i-1]
					tmpa[len(tmpa)-i-1] = tmp
				}
			}
		}
		arr = append(arr, tmpa)
		nodes = tmp
		level++
	}

	return arr
}
