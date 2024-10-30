/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeQueries(root *TreeNode, queries []int) []int {
    h := map[*TreeNode]int{}


    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        l := dfs(root.Left)
        r := dfs(root.Right)
        h[root] = max(l, r) + 1
        return h[root]
    }
    dfs(root)

    ans := make([]int, len(h) + 1)
    var dfs1 func(*TreeNode, int, int)
    dfs1 = func(node *TreeNode, d, rest int) {
        if node == nil {
            return
        }
        d++
        ans[node.Val] = rest
        dfs1(node.Left, d, max(rest, d + h[node.Right]))
        dfs1(node.Right, d, max(rest, d + h[node.Left]))
    }
    dfs1(root, -1, 0)

    for i, x := range queries {
        queries[i] = ans[x]
    }

    return queries
}