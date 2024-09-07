/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    } else if root.Left == nil {
        return root
    }

    
    tmp := upsideDownBinaryTree(root.Left)

    root.Left.Right = root
    root.Left.Left = root.Right
    root.Left = nil
    root.Right = nil

    return tmp
}