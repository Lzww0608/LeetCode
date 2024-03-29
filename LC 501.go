/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) (ans []int) {
    pre, cnt, mx := math.MaxInt32, 0, 0

    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        if pre == root.Val {
            cnt++
        } else {
            pre, cnt = root.Val, 1
        }
        if cnt == mx {
            ans = append(ans, pre)
        } else if cnt > mx {
            mx = cnt
            ans = []int{pre}
        }
        dfs(root.Right)
        
    }
    dfs(root)

    
    return
}