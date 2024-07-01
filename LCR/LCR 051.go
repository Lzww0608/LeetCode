/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := math.MinInt

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        cur := max(root.Val, root.Val + l, root.Val + r)
        ans = max(ans, cur, root.Val + l + r)
        return cur
    }
    dfs(root)

    return ans
}