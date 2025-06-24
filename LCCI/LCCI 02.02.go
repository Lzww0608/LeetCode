/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    p, q := head, head 
    for i := 0; i < k; i++ {
        q = q.Next
    }
    for q != nil {
        p, q = p.Next, q.Next
    }

    return p.Val
}