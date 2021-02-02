package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
      if root == nil {
            return nil
      }
      arr := [][]int{}
      nodes := []*TreeNode{root}
      tmpa := []int{}
      for len(nodes) >0 {
            tmpa = []int{}
            tmp := []*TreeNode{}
            for i := 0; i < len(nodes); i++ {
                  tmpa = append(tmpa, nodes[i].Val)
                  if nodes[i].Left != nil {
                        tmp = append(tmp, nodes[i].Left)
                  }
                  if nodes[i].Right != nil {
                        tmp = append(tmp, nodes[i].Right)
                  }
            }
            arr = append(arr, tmpa)
            nodes = tmp
      }

      return arr
}
