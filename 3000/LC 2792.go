/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countGreatEnoughNodes(root *TreeNode, k int) (ans int) {
    var dfs func(*TreeNode) []int
    dfs = func(root *TreeNode) []int {
        if root == nil {
            return nil
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        x := root.Val
        cur := []int{x}
        cur = append(cur, l...)
        cur = append(cur, r...)
        sort.Ints(cur)
        cur = cur[:min(len(cur), k)]
        if len(cur) == k && cur[k-1] < x {
            ans++
        }

        return cur
    }
    dfs(root)

    return
}