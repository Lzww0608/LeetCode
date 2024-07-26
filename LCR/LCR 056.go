/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    q := []*TreeNode{root}
    m := map[int]struct{}{}

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if _, ok := m[k - node.Val]; ok {
            return true
        }
        m[node.Val] = struct{}{}

        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }

    return false
}