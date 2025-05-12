/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidSequence(root *TreeNode, arr []int) bool {
    if root == nil || len(arr) == 0 || root.Val != arr[0] {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return len(arr) == 1
    }

    return isValidSequence(root.Left, arr[1:]) || isValidSequence(root.Right, arr[1:])
}