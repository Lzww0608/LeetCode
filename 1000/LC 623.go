/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth == 1 {
        ans := &TreeNode{val, root ,nil}
        return ans
    }

    q := []*TreeNode{root}
    for depth > 1 {
        if depth == 2 {
            for _, v := range q {
                l := &TreeNode{val, v.Left, nil}
                r := &TreeNode{val, nil, v.Right}
                v.Left = l
                v.Right = r
            }
            break
        }

        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            if cur.Left != nil {
                q = append(q, cur.Left)
            }
            if cur.Right != nil {
                q = append(q, cur.Right)
            }
        }
        depth--
    }

    return root
}