/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimalExecTime(root *TreeNode) float64 {
    _, t := dfs(root)
    return t
}

func dfs(root *TreeNode) (s, t float64) {
    if root == nil {
        return 0.0, 0.0
    }

    s1, t1 := dfs(root.Left)
    s2, t2 := dfs(root.Right)

    return s1 + s2 + float64(root.Val), max(t1, t2, (s1 + s2) / 2) + float64(root.Val)
}