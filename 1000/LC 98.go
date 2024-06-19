/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return preOrder(root, math.MinInt, math.MaxInt)
}

func preOrder(root *TreeNode, left, right int) bool {
    if root == nil {
        return true
    }
    x := root.Val
    return x > left && x < right && preOrder(root.Left, left, x) && preOrder(root.Right, x, right)
}