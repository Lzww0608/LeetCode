/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) (ans int) {
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return 
        }
        sum = sum * 10 + root.Val
        if root.Left == nil && root.Right == nil {
            ans += sum
            return
        }
        dfs(root.Left, sum)
        dfs(root.Right, sum)
    }
    dfs(root, 0)
    return
}
