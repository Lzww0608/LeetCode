/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    ans := []int{}
    q := []*TreeNode{}
    if root != nil {
        q = append(q, root)
    }
    for len(q) > 0 {
        maxVal := math.MinInt32
        for k := len(q); k > 0; k-- {
            t := q[0]
            q = q[1:]
            maxVal = max(maxVal, t.Val)
            if t.Left != nil {
                q = append(q, t.Left)
            }
            if t.Right != nil {
                q = append(q, t.Right)
            }
        }
        ans = append(ans, maxVal)
    }
    return ans
}