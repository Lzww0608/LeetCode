/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
        return nil
    }
    l := trimBST(root.Left, low, high)
    r := trimBST(root.Right, low, high)
    if root.Val > high || root.Val < low {
        if l != nil {
            return l
        }
        return r
    }
    root.Left = l 
    root.Right = r
    return root
}