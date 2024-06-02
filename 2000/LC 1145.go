/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    sum, node := n, find(root, x)
    sum_x, sum_l, sum_r := cnt(node), cnt(node.Left), cnt(node.Right)
    if sum - sum_x > sum_x || sum - sum_l < sum_l || sum - sum_r < sum_r{
        return true
    }
    return false
}

func find(head *TreeNode, x int) *TreeNode {
    if head == nil {
        return nil
    } else if head.Val == x {
        return head
    }
    l := find(head.Left, x)
    if l == nil {
        return find(head.Right, x)
    } 
    return l
}

func cnt(head *TreeNode) int {
    if head == nil {
        return 0
    }
    return 1 + cnt(head.Left) + cnt(head.Right)
}
