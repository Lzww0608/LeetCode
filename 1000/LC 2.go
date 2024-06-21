/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(p *ListNode, q *ListNode) *ListNode {
    dummy := &ListNode{-1, nil}
    tmp := dummy
    add := 0
    for p != nil || q != nil || add > 0 {
        t := add
        if p != nil {
            t += p.Val
            p = p.Next
        }
        if q != nil {
            t += q.Val
            q = q.Next
        }
        add = t / 10
        tmp.Next = &ListNode{t % 10, nil}
        tmp = tmp.Next
    }

    return dummy.Next
}
