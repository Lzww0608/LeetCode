/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) (ans int) {
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, mx int) {
        if root == nil {
            return
        }
        if root.Val >= mx {
            ans++
            mx = root.Val
        }
        dfs(root.Left, mx)
        dfs(root.Right, mx)
    }
    dfs(root, math.MinInt32)

    return
}