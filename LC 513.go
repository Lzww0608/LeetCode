/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    ans, depth := root.Val, 0
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, d int) {
        if root == nil {
            return 
        }
        dfs(root.Left, d + 1)
        dfs(root.Right, d + 1)
        if root.Left == nil && root.Right == nil && d > depth {
            depth = d 
            ans = root.Val
        }
    }
    dfs(root, 0)
    return ans
}