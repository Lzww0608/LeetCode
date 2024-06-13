/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) (ans int) {
    m := map[int]int{}
    m[0] = 1
    
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return
        }
        sum += root.Val
        ans += m[sum - targetSum]
        m[sum]++
        dfs(root.Left, sum)
        dfs(root.Right, sum)
        m[sum]--
    }
    dfs(root, 0)

    return
}