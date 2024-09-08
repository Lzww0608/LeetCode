/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const N int = 101
func findLeaves(root *TreeNode) (ans [][]int) {

    var dfs func(*TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        l := dfs(root.Left)
        r := dfs(root.Right)
        d := max(l, r)
        if d >= len(ans) {
            ans = append(ans, []int{})
        }
        ans[d] = append(ans[d], root.Val)
        return d + 1
    }

    dfs(root)

    return ans
}