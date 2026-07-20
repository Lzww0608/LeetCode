/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countDominantNodes(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) int 
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        l, r := dfs(node.Left), dfs(node.Right)
        if node.Val >= max(l, r) {
            ans++
        }
        return max(node.Val, l, r)
    }

    dfs(root)
    return
}