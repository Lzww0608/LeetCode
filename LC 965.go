/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    val := root.Val
    q := []*TreeNode{root}
    for len(q) > 0 {
        t := q[0]
        q = q[1:]
        if t.Val != val {
            return false
        }
        if t.Left != nil {
            q = append(q, t.Left)
        }
        if t.Right != nil {
            q = append(q, t.Right)
        }
    }

    return true
}