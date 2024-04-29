/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
    path := []int{}

    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return 
        }
        sum += root.Val
        path = append(path, root.Val)
        if root.Left == nil && root.Right == nil {
            if sum == targetSum {
                ans = append(ans, append([]int(nil), path...)) 
            }
        } else {
            dfs(root.Left, sum)
            dfs(root.Right, sum)
        }
        sum -= root.Val
        path = path[:len(path)-1]
    }

    dfs(root, 0)

    return 
}