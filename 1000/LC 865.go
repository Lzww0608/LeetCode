/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) (ans *TreeNode) {
    mx := -1

    var dfs func(*TreeNode, int) int
    dfs = func(root *TreeNode, d int) int {
        if root == nil {
            mx = max(mx, d)
            return d
        }
        l := dfs(root.Left, d + 1)
        r := dfs(root.Right, d + 1)
        if l == r && l == mx {
            ans = root
        }
        return max(l, r)
    }
    dfs(root, 0)

    return
}