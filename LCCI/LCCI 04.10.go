/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkSubTree(p *TreeNode, q *TreeNode) bool {
    if p == nil {
        return false
    } else if p.Val == q.Val && isSub(p, q) || q == nil {
        return true
    }

    return checkSubTree(p.Left, q) || checkSubTree(p.Right, q)
}

func isSub(p, q *TreeNode) bool {
    if p == q {
        return true
    } else if p == nil || q == nil || p.Val != q.Val {
        return false
    }

    return isSub(p.Left, q.Left) && isSub(p.Right, q.Right)
}
