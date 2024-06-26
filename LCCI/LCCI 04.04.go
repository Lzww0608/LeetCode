/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) (ans bool) {
    ans = true
    
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil || ans == false {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        if abs(l - r) > 1 {
            ans = false
        }
        return max(l, r) + 1
    }
    dfs(root)

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}