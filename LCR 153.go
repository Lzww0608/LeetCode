/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathTarget(root *TreeNode, target int) (ans [][]int) {
    path := []int{}
    sum := 0
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        path = append(path, root.Val)
        sum += root.Val
        if root.Left == nil && root.Right == nil && sum == target {
            ans = append(ans, append([]int(nil), path...))
            sum -= root.Val
            path = path[:len(path)-1]
            return
        }
        
        if root.Left != nil {
            dfs(root.Left)
        }
        if root.Right != nil {
            dfs(root.Right)
        }
        sum -= root.Val
        path = path[:len(path)-1]
    }

    if root != nil {
        dfs(root)
    }

    return
}