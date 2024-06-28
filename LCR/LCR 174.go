/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTargetNode(root *TreeNode, cnt int) int {
    ans, d := math.MinInt32, 0
    
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil || ans != math.MinInt32 {
            return
        }
        
        dfs(root.Right)
        d++
        if d == cnt {
            ans = root.Val
            return
        }
        dfs(root.Left)
        
    }
    dfs(root)

    return ans
}
