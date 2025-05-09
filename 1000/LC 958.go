/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
    q := []*TreeNode{root}

    isEnd := false
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        if isEnd && cur != nil {
            return false
        }
        if cur == nil {
            isEnd = true
        } else {
            q = append(q, cur.Left)
            q = append(q, cur.Right)
        }
    }

    return true
}