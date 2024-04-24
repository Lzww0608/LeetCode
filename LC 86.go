/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    p := &ListNode{-1, nil}
    q := &ListNode{-1, nil}

    pp, qq := p, q
    for head != nil {
        if head.Val < x {
            pp.Next = head
            pp = pp.Next
        } else {
            qq.Next = head
            qq = qq.Next
        }
        head = head.Next
    }
    qq.Next = nil // ***
    pp.Next = q.Next

    return p.Next
}