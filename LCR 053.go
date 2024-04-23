/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
    cur := root

    for cur != nil {
        if cur.Val > p.Val {
            ans = cur
            cur = cur.Left
        } else {
            cur = cur.Right
        }
    }

    return
}