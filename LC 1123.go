/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) (ans *TreeNode) {
    md := -1
    var dfs func(*TreeNode, int) int
    dfs = func(root *TreeNode, d int) int {
        if root == nil {
            md = max(md, d)
            return d
        }
        l, r := dfs(root.Left, d + 1), dfs(root.Right, d + 1)
        if l == r && l == md {
            ans = root
        }
        return max(l, r)
    }
    dfs(root, 0)
    return 
}