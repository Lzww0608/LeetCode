/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    b := true

    var dfs func(*TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil || b == false {
            return 0
        }
        l, r := dfs(root.Left), dfs(root.Right)
        if abs(l - r) > 1 {
            b = false 
            return 0 
        }
        return max(l, r) + 1
    }
    dfs(root)

    return b
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}
