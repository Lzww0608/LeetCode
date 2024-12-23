/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
    for root != nil {
        if root.Val <= p.Val {
            root = root.Right
        } else {
            ans = root
            root = root.Left
        }
    }

    return
}