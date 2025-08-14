/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    dfs(root.Left, root.Right, 0)

    return root
}

func dfs(l, r *TreeNode, k int) {
    if l == nil || r == nil {
        return
    }

    if k == 0 {
        l.Val, r.Val = r.Val, l.Val
    }
    dfs(l.Left, r.Right, k ^ 1)
    dfs(l.Right, r.Left, k ^ 1)
}