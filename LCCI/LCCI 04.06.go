/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
    target := p.Val
    for root != nil {
        if root.Val > target {
            ans = root
            root = root.Left
        } else {
            root = root.Right
        }
    }

    return
}