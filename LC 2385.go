/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) (ans int) {
    h := -1

    var dfs func(*TreeNode, int) int
    dfs = func(root *TreeNode, d int) int {
        if root == nil {
            return 0
        }
        
        if root.Val == start {
            h = d
        }

        l := dfs(root.Left, d + 1)
        in := false
        if h != -1 {
            in = true
        }
        r := dfs(root.Right, d + 1)

        if root.Val == start {
            ans = max(ans, l, r)
        } else if in {
            ans = max(ans, h - d + r)
        } else {
            ans = max(ans, h - d + l)
        }

        return max(l, r) + 1
    }
    dfs(root, 0)

    return
}