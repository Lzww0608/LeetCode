/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumFlips(root *TreeNode, result bool) int {
    a, b := dfs(root)
    if result {
        return a
    }
    return b
}

func dfs(root *TreeNode) (int, int) {
    if root.Val < 2 {
        if root.Val == 0 {
            return 1, 0
        } else {
            return 0, 1
        }
    } else if root.Val == 5 {
        var l, r int
        if root.Left != nil {
            l, r = dfs(root.Left)
        } else {
            l, r = dfs(root.Right)
        }
        return r, l
    }

    l1, l2 := dfs(root.Left)
    r1, r2 := dfs(root.Right)

    if root.Val == 2 {
        return min(l1 + r1, l1 + r2, l2 + r1), l2 + r2
    } else if root.Val == 3 {
        return l1 + r1, min(l1 + r2, l2 + r1, l2 + r2)
    } else {
        return min(l1 + r2, l2 + r1), min(l1 + r1, l2 + r2)
    }
}