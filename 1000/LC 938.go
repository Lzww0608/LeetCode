/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        if root.Val >= low && root.Val <= high {
            sum += root.Val
            dfs(root.Right)
            dfs(root.Left)
        } else if root.Val > high {
            dfs(root.Left)
        } else {
            dfs(root.Right)
        }
    }
    dfs(root)
    return sum
}
