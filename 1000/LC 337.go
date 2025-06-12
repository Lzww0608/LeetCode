/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) (int, int)
    dfs = func(root *TreeNode) (a, b int) {
        if root == nil {
            return
        }
        la, lb := dfs(root.Left)
        ra, rb := dfs(root.Right)
        a = max(root.Val + lb + rb)
        b = max(la, lb) + max(ra, rb)
        ans = max(ans, a, b)
        return
    }

    return max(dfs(root))
}