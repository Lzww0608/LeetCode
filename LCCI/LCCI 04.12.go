/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) (ans int) {
    m := map[int]int{0: 1}

    var dfs func(*TreeNode, int) 
    dfs = func(root *TreeNode, pre int) {
        if root == nil {
            return
        }

        pre += root.Val
        ans += m[pre-sum]
        m[pre]++
        dfs(root.Left, pre)
        dfs(root.Right, pre)
        m[pre]--
    }
    dfs(root, 0)

    return
}