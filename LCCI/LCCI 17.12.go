/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBiNode(root *TreeNode) *TreeNode {
    var pre, head *TreeNode

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if pre == nil {
            head = root
        } else {
            pre.Right = root
        }
        root.Left = nil
        pre = root
        dfs(root.Right)
        
    }
    dfs(root)

    return head
}