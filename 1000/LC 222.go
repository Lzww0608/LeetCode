/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    l, r := 0, 0

    for p := root; p != nil; p = p.Left {
        l++
    }
    for p := root; p != nil; p = p.Right {
        r++
    }

    if l == r {
        return quickPow(2, l) - 1
    }

    return countNodes(root.Left) + countNodes(root.Right) + 1
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res *= a
        }
        a *= a
        r >>= 1
    }

    return res
}