/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    ans, pre := math.MaxInt32, -1

    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if pre != -1 {
            ans = min(ans, root.Val - pre)
        }
        pre = root.Val
        dfs(root.Right)
    }
    dfs(root)

    return ans
}