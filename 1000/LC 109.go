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
func sortedListToBST(head *ListNode) *TreeNode {
    cur, n := head, 0
    for cur != nil {
        n++
        cur = cur.Next
    }

    var dfs func(int, int) *TreeNode 
    dfs = func(l, r int) *TreeNode {
        if l > r {
            return nil
        }
        mid := l + ((r - l) >> 1)
        left := dfs(l, mid - 1)
        node := &TreeNode{head.Val, left, nil}
        head = head.Next
        node.Right = dfs(mid + 1, r)
        return node
    }

    return dfs(0, n - 1)
}
