/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
    if root1 == nil {
        return false
    }

    return myFind(root2, target - root1.Val) || twoSumBSTs(root1.Left, root2, target) || twoSumBSTs(root1.Right, root2, target)
}

func myFind(root *TreeNode, target int) bool {
    for root != nil {
        if root.Val == target {
            return true
        } else if root.Val < target {
            root = root.Right
        } else {
            root = root.Left
        }
    }

    return false
}