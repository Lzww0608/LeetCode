/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkSymmetricTree(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return match(root.Left, root.Right)
}

func match(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil {
        return false
    } else if p.Val != q.Val {
        return false
    }

    return match(p.Right, q.Left) && match(p.Left, q.Right)
}
