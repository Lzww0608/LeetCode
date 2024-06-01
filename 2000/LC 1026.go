/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) (ans int) {
    var dfs func(root *TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        mn, mx := root.Val, root.Val
        if root.Left != nil {
            a, b := dfs(root.Left)
            mn = min(mn, a)
            mx = max(mx, b)
        }
        if root.Right != nil {
            a, b := dfs(root.Right)
            mn = min(mn, a)
            mx = max(mx, b)
        }
        ans = max(ans, mx - root.Val, root.Val - mn)
        return mn, mx
    }
    dfs(root)

    return 
}
