/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    r := root.Right
    root.Right = mirrorTree(root.Left)
    root.Left = mirrorTree(r)

    return root
}
