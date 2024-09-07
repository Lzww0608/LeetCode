/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"
func closestKValues(root *TreeNode, target float64, k int) (ans []int) {

    var dfs func(*TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }

        dfs(root.Left)

        if len(ans) < k {
            ans = append(ans, root.Val)
        } else {
            if math.Abs(float64(ans[0]) - target) > math.Abs(float64(root.Val) - target) {
                ans = ans[1:]
                ans = append(ans, root.Val)
            } else {
                return
            }
        }

        dfs(root.Right)
    }
    dfs(root)

    return 
}