/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if root == nil {
        return nil
    }
    limit -= root.Val
    if root.Left == root.Right {
        if limit > 0 {
            return nil
        }
        return root
    }
    root.Left, root.Right = sufficientSubset(root.Left, limit), sufficientSubset(root.Right, limit)
    if root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}
