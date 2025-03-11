/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func equalToDescendants(root *TreeNode) (ans int) {
    
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) (sum int) {
        if node == nil {
            return 0
        }
        sum += dfs(node.Left) + dfs(node.Right)
        if sum == node.Val {
            ans++
        }
        return sum + node.Val
    }
    dfs(root)

    return
}