/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{val, nil, nil}
    }
    res := root
    for {
        if root.Val > val {
            if root.Left == nil {
                root.Left = &TreeNode{val, nil, nil}
                break
            }
            root = root.Left
        } else {
            if root.Right == nil {
                root.Right = &TreeNode{val, nil, nil}
                break
            }
            root = root.Right
        }
    }

    return res
}