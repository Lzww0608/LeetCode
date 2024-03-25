/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
    var check func(*TreeNode, *ListNode) bool
    check = func(root *TreeNode, p *ListNode) bool {
        if p == nil {
            return true
        } else if root == nil || root.Val != p.Val {
            return false
        }
        return check(root.Left, p.Next) || check(root.Right, p.Next)
    }


    var dfs func(*TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return false
        }
        if root.Val == head.Val {
            //p, q := head, root
            if check(root, head) {
                return true
            }
        } 
        return dfs(root.Left) || dfs(root.Right)
    }
    return dfs(root)
}