

import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32
    
    var dfs func(*TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        ans = max(ans, l + r + root.Val)
        return max(0, l + root.Val, r + root.Val)
    }
    dfs(root)

    return ans
}