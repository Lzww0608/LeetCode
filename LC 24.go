/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    p := dummy
    for head != nil && head.Next != nil {
        q := head.Next
        next := q.Next
        p.Next, q.Next = q, head
        head, p = next, head
    }
    p.Next = head

    return dummy.Next
}