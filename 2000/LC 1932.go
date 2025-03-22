

import "math"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func canMerge(trees []*TreeNode) *TreeNode {
    leaves := make(map[int]bool)
    candidate := make(map[int]*TreeNode)

    for _, v := range trees {
        if v.Left != nil {
            leaves[v.Left.Val] = true
        } 
        if v.Right != nil {
            leaves[v.Right.Val] = true
        }
        candidate[v.Val] = v
    }

    prev := math.MinInt32
    var dfs func(*TreeNode) bool 
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return true
        }

        if root.Left == nil && root.Right == nil && candidate[root.Val] != nil {
            root.Left = candidate[root.Val].Left
            root.Right = candidate[root.Val].Right
            delete(candidate, root.Val)
        }

        if !dfs(root.Left) {
            return false
        }

        if root.Val <= prev {
            return false
        }
        prev = root.Val
        return dfs(root.Right) 
    }

    for _, t := range trees {
        if !leaves[t.Val] {
            delete(candidate, t.Val)
            if dfs(t) && len(candidate) == 0 {
                return t 
            }

            return nil
        }
    }

    return nil
}