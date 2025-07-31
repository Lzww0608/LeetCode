/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) (ans int) {
    var dfs func(root *TreeNode) (int, int)
    dfs = func(root *TreeNode) (val, cnt int) {
        if root == nil {
            return -1, 0
        }

        lVal, lCnt := dfs(root.Left)
        rVal, rCnt := dfs(root.Right)
        val, cnt = root.Val, 1 
        cur := 1
        if lVal == val {
            cur += lCnt
            cnt = max(cnt, 1 + lCnt)
        }
        if rVal == val {
            cur += rCnt
            cnt = max(cnt, 1 + rCnt)
        }
        ans = max(ans, cur - 1)
        return 
    }
    dfs(root)

    return
}