/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    m := map[int]bool{}
    q := []*TreeNode{root}

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if m[k - node.Val] {
            return true
        }
        m[node.Val] = true
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }

    return false
}