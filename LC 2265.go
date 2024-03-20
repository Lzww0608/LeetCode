/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) (ans int) {

    var dfs func(*TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        sum, cnt := root.Val, 1

        if root.Left != nil {
            a, b := dfs(root.Left)
            sum += a
            cnt += b
        } 
        if root.Right != nil {
            a, b := dfs(root.Right)
            sum += a
            cnt += b
        }

        if sum / cnt == root.Val {
            ans++
        }

        return sum, cnt
    }

    dfs(root)
    return
}