
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) (ans int) {

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        res := 1
        if root.Left == nil {
            res = max(res, 1)
        } else if root.Left.Val - 1 == root.Val {
            res = max(res, l + 1)
        }

        if root.Right == nil {
            res = max(res, 1)
        } else if root.Right.Val - 1 == root.Val {
            res = max(res, r + 1)
        }
        ans = max(ans, res)

        return res
    }
    dfs(root)

    return 
}