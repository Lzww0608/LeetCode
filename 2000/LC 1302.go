/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) (ans int) {
    mx := 0
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, d int) {
        if root == nil {
            return
        }
        
        if d > mx {
            ans = root.Val
            mx = d
        } else if d == mx {
            ans += root.Val
        }
        dfs(root.Left, d + 1)
        dfs(root.Right, d + 1)
    }
    dfs(root, 0)

    return
}