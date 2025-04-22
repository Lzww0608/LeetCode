/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
    q := []*TreeNode{root}

    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            if cur == u {
                if sz > 1 {
                    return q[0]
                }
                return nil
            }
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
    }

    return nil
}