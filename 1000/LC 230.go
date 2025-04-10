/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) (ans int) {
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        dfs(root.Left)
        if k--; k == 0 {
            ans = root.Val
        }
        dfs(root.Right)
        return
    }
    
    dfs(root)
    return
}