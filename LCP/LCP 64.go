/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type pair struct {
    node *TreeNode
    a, b bool
}
var memo map[pair]int
func closeLampInTree(root *TreeNode) int {
    memo = make(map[pair]int)
    return dfs(root, false, false)
}

func dfs(root *TreeNode, a, b bool) int {
    if root == nil {
        return 0
    }
    if v, ok := memo[pair{root, a, b}]; ok {
        return v
    }

    cur := root.Val
    if a {
        cur ^= 1
    }
    if b {
        cur ^= 1
    }

    res := math.MaxInt32
    if cur == 0 {
        // 1 2, 1 3, 2 3
        res = min(res, dfs(root.Left, a, false) + dfs(root.Right, a, false))
        res = min(res, dfs(root.Left, a, true) + dfs(root.Right, a, true) + 2)
        res = min(res, dfs(root.Left, !a, false) + dfs(root.Right, !a, false) + 2)      
        res = min(res, dfs(root.Left, !a, true) + dfs(root.Right, !a, true) + 2)
    } else {
        // 1, 2, 3, 1 2 3
        res = min(res, dfs(root.Left, a, false) + dfs(root.Right, a, false) + 1)
        res = min(res, dfs(root.Left, a, true) + dfs(root.Right, a, true) + 1)
        res = min(res, dfs(root.Left, !a, false) + dfs(root.Right, !a, false) + 1)
        res = min(res, dfs(root.Left, !a, true) + dfs(root.Right, !a, true) + 3)
    }
    memo[pair{root, a, b}] = res

    return res
} 