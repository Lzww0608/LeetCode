/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) (ans []int) {
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        if root.Left != nil && root.Right == nil {
            ans = append(ans, root.Left.Val)
        }
        if root.Left == nil && root.Right != nil {
            ans = append(ans, root.Right.Val)
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)

    return 
}