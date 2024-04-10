/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) (ans int) {
    q := []*TreeNode{root}

    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node.Left == nil && node.Right == nil {
            ans += node.Val
            continue
        }
        if node.Left != nil {
            cur := (node.Val << 1) | node.Left.Val
            node.Left.Val = cur
            q = append(q, node.Left)  
        }
        if node.Right != nil {
            cur := (node.Val << 1) | node.Right.Val
            node.Right.Val = cur
            q = append(q, node.Right)
        }
    }

    return 
}