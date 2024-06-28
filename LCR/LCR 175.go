/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func calculateDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := calculateDepth(root.Left)
    r := calculateDepth(root.Right)
    return max(l, r) + 1
}
