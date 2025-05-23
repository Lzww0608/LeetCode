/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func splitBST(root *TreeNode, target int) (ans []*TreeNode) {
    if root == nil {
        return []*TreeNode{nil, nil}
    } else if root.Val <= target {
        ans = splitBST(root.Right, target)
        root.Right = ans[0]
        ans[0] = root
    } else {
        ans = splitBST(root.Left, target)
        root.Left = ans[1]
        ans[1] = root
    }

    return
}