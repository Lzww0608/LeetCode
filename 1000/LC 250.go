

import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countUnivalSubtrees(root *TreeNode) (ans int) {
    
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return math.MinInt32
        }

        l := dfs(root.Left)
        r := dfs(root.Right)
        if l == math.MinInt32 {
            l = root.Val
        }
        if r == math.MinInt32 {
            r = root.Val
        }
        if (l == r && l == root.Val)  {
            ans++
            return root.Val
        }
        
        return math.MaxInt32
    }

    dfs(root)
    return
}