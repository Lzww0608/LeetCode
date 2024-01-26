/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    ans, q := []int{}, []*TreeNode{}
    if root != nil {
        q = append(q, root)
    }
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
            if t == 1 {
                ans = append(ans, node.Val)
            }
        }
    }
    return ans
}