/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
    layer := 0
    if root == nil {
        return 
    }
    q := []*TreeNode{root}
    for len(q) > 0 {
        layer++
        level := []int{}
        for k := len(q); k > 0; k-- {
            node := q[0]
            q = q[1:]
            if layer & 1 == 1 {
                level = append(level, node.Val)
            } else {
                level = append([]int{node.Val}, level...)
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        ans = append(ans, append([]int(nil), level...))
    }
    return 
}