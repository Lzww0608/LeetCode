/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
    q := []*TreeNode{root}

    d := 0
    for len(q) > 0 {
        pre := -1
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            if cur.Val & 1 == d {
                return false
            }
            if pre != -1 {
                if d == 1 && cur.Val >= pre || d == 0 && cur.Val <= pre {
                    return false
                }
            }
            pre = cur.Val
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        d ^= 1
    }

    return true
}