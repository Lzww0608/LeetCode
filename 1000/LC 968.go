/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) (ans int) {
    //0：无摄像头无覆盖
    //1：有摄像头有覆盖
    //2：无摄像头有覆盖
    var dfs func(*TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 2
        }
        l, r := dfs(root.Left), dfs(root.Right)
        
        if l == 0 || r == 0 {
            ans++
            return 1
        } else if l == 1 || r == 1 {
            return 2
        }
        return 0
    }

    if dfs(root) == 0 {
        ans++
    }

    return 
}
