/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(p *ListNode, q *ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    head := dummy

    for p != nil && q != nil {
        if p.Val <= q.Val {
            head.Next = p
            head, p = head.Next, p.Next
        } else {
            head.Next = q
            head, q = head.Next, q.Next
        }
    }

    if p != nil {
        head.Next = p
    } else {
        head.Next = q
    }

    return dummy.Next
}