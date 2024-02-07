/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    depth_x, depth_y := -1, -1
    fa_x, fa_y := root, root
    q := []*TreeNode{root}
    d := 0
    for len(q) > 0 && (depth_x == -1 || depth_y == -1) {
        d++
        for k := len(q); k > 0; k-- {
            t := q[0]
            q = q[1:]
            if t.Left != nil {
                if t.Left.Val == x {
                    depth_x = d
                    fa_x = t
                } else if t.Left.Val == y {
                    depth_y = d
                    fa_y = t
                }
                q = append(q, t.Left)
            }
            if t.Right != nil {
                if t.Right.Val == x {
                    depth_x = d
                    fa_x = t
                } else if t.Right.Val == y {
                    depth_y = d
                    fa_y = t
                }
                q = append(q, t.Right)
            }
        }
    }
    return fa_x != fa_y && depth_x == depth_y
}