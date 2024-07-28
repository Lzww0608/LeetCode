/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    dummy := &TreeNode{}
    p := dummy

    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        p.Right = root
        root.Left = nil
        p = p.Right
        dfs(root.Right)
    }
    dfs(root)

    return dummy.Right
}