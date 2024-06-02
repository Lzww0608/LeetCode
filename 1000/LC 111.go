/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if root.Left == nil || root.Right == nil {
        return l + r + 1
    }

    return min(l, r) + 1
}
