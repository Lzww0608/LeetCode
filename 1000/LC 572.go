/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    q := []*TreeNode{root}

    var check func(*TreeNode, *TreeNode) bool
    check = func(p, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        } else if p == nil || q == nil || p.Val != q.Val {
            return false
        }
        return check(p.Left, q.Left) && check(p.Right, q.Right)
    }

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node.Val == subRoot.Val && check(node, subRoot) {
            return true
        }
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }
    return false
}
