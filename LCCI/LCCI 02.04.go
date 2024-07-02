/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    sml_dummy := &ListNode{-1, nil}
    big_dummy := &ListNode{-1, nil}
    p, q := sml_dummy, big_dummy
    for head != nil {
        if head.Val < x {
            p.Next = head
            p = p.Next
        } else {
            q.Next = head
            q = q.Next
        }
        head = head.Next
    }

    p.Next = big_dummy.Next
    q.Next = nil
    return sml_dummy.Next
}