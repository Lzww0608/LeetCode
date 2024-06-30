/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    ans, mn := -1, root.Val
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil || ans != -1 && root.Val >= ans {
            return
        }
        
        if root.Val > mn {
            ans = root.Val
        }

        dfs(root.Left)
        dfs(root.Right)
        
    }
    dfs(root)

    return ans
}
