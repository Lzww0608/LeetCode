/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) (ans int) {

    var dfs func(*TreeNode) []int
    dfs = func(root *TreeNode) []int {
        if root == nil {
            return nil
        } else if root.Left == nil && root.Right == nil {
            return []int{0}
        }

        res := []int{}
        for _, d := range dfs(root.Left) {
            if d + 1 < distance {
                res = append(res, d + 1)
            }
        }
        n := len(res)
        for _, d := range dfs(root.Right) {
            if d + 1 < distance {
                for i := 0; i < n; i++ {
                    if res[i] + d + 1 <= distance {
                        ans++
                    }
                }
                res = append(res, d + 1)
            }
        }

        return res
    }
    dfs(root)

    return
}