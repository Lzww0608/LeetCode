/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    ans := true

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil || ans == false {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        if l - r > 1 || r - l > 1 {
            ans = false
            return 0
        } 
        return max(l, r) + 1
    }
    dfs(root)

    return ans
}
