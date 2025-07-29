/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func heightOfTree(root *TreeNode) int {
    if root == nil || root.Left == nil && root.Right == nil {
        return 0
    }

    if root.Left != nil && root.Left.Right == root || root.Right != nil && root.Right.Left == root {
        return 0
    }

    return max(heightOfTree(root.Left), heightOfTree(root.Right)) + 1
}