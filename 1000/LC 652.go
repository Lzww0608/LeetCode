/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) (ans []*TreeNode) {
    cnt := make(map[string]int)

    var dfs func(*TreeNode) string 
    dfs = func(root *TreeNode) string {
        if root == nil {
            return "#"
        }

        l := dfs(root.Left)
        r := dfs(root.Right)
        cur := strconv.Itoa(root.Val) + "," + l + "," + r 
        if cnt[cur] == 1 {
            ans = append(ans, root)
        }
        cnt[cur]++
        return cur
    }
    dfs(root)

    return
}