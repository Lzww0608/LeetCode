/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil || B == nil {
        return false
    }
    return match(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func match(A, B *TreeNode) bool {
    if B == nil {
        return true
    } else if A == nil {
        return false
    } else if A.Val != B.Val {
        return false
    }

    return match(A.Left, B.Left) && match(A.Right, B.Right)
}
