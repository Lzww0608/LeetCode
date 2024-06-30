/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    pre, p := dummy, head
    for p != nil && p.Next != nil {
        if p.Val == p.Next.Val {
            t := p.Val
            for p != nil && p.Val == t {
                p = p.Next
            }
            pre.Next = p
        } else {
            pre, p = p, p.Next
        }
    }
    return dummy.Next
}
