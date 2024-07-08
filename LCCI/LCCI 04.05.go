/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    ans := true

    prev := math.MinInt
    var dfs func(*TreeNode) 
    dfs = func(root *TreeNode) {
        if !ans || root == nil {
            return
        }
        dfs(root.Left)
        if prev >= root.Val {
            ans = false
            return
        }
        prev = root.Val
        dfs(root.Right)
    }
    dfs(root)

    return ans
}