/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    ans := 0
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        d := dfs(root.Left) + dfs(root.Right) + root.Val - 1
        ans += abs(d)
        return d
    }
    dfs(root)
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}