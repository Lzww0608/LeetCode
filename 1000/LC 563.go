/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    ans := 0
    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        a := dfs(root.Left)
        b := dfs(root.Right)
        ans += abs(a - b)
        return a + b + root.Val
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
