/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) (ans []int) {
    i := 0

    var dfs func(*TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return true
        } else if root.Val != voyage[i] {
            return false
        }
        i++
        if root.Left != nil && root.Left.Val != voyage[i] {
            ans = append(ans, root.Val)
            return dfs(root.Right) && dfs(root.Left)
        }
        return dfs(root.Left) && dfs(root.Right)
    }

    if dfs(root) {
        return 
    }

    return []int{-1}
}