/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) (ans int) {
    var dfs func(*TreeNode) (int, int, int)
    dfs = func(root *TreeNode) (int, int, int) {
        if root == nil {
            return math.MaxInt32, math.MinInt32, 0
        }
        
        lMin, lMax, lSum := dfs(root.Left)
        rMin, rMax, rSum := dfs(root.Right)
        if root.Val <= lMax || root.Val >= rMin {
            return math.MinInt32, math.MaxInt32, 0
        }

        sum := lSum + rSum + root.Val
        ans = max(ans, sum)

        return min(root.Val, lMin), max(root.Val, rMax), sum
    }
    dfs(root)

    return
}