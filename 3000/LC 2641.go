/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    q := []*TreeNode{root}

    for len(q) > 0 {
        tmp := q
        q = nil
        sum := 0
        
        for _, node := range tmp {
            if node.Left != nil {
                q = append(q, node.Left)
                sum += node.Left.Val
            }
            if node.Right != nil {
                q = append(q, node.Right)
                sum += node.Right.Val
            }
        }

        for _, node := range tmp {
            son := 0
            if node.Left != nil {
                son += node.Left.Val
            }
            if node.Right != nil {
                son += node.Right.Val
                node.Right.Val = sum - son
            }
            if node.Left != nil {
                node.Left.Val = sum - son
            }
        }
    }

    root.Val = 0
    return root
}