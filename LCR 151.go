/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func decorateRecord(root *TreeNode) (ans [][]int) {
    if root == nil {
        return
    }
    d := 0
    q := []*TreeNode{root}
    for len(q) > 0 {
        n := len(q)
        tmp := make([]int, n)
        for k := 0; k < n; k++ {
            node := q[0]
            q = q[1:]
            if d & 1 == 0 {
                tmp[k] = node.Val
            } else {
                tmp[n-k-1] = node.Val
            }

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        d++
        ans = append(ans, tmp)
    }

    return 
}