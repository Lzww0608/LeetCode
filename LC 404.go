/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) (ans int) {
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
            ans += root.Left.Val
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return
}