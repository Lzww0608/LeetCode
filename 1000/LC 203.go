/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{-1, head}
    pre, p := dummy, head

    for p != nil {
        if p.Val == val {
            p = p.Next
            pre.Next = p
        } else {
            pre, p = p, p.Next
        }
    }

    return dummy.Next
}
