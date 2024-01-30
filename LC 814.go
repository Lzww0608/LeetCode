/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    var dfs func(root *TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return true
        }
        f1, f2 := dfs(root.Left), dfs(root.Right)
        if f1 {
            root.Left = nil
        }
        if f2 {
            root.Right = nil
        }
        return f1 && f2 && root.Val == 0
    }
    if dfs(root) {
        return nil
    }
    return root
}