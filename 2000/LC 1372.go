/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) (ans int) {
    
    var dfs func(*TreeNode, int, int)
    dfs = func(root *TreeNode, l, r int) {
        ans = max(ans, l, r)
        if root.Left != nil {
            dfs(root.Left, r + 1, 0)
        }
        if root.Right != nil {
            dfs(root.Right, 0, l + 1)
        }
    }
    dfs(root, 0, 0)

    return 
}
