/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(a *ListNode, b *ListNode) *ListNode {
    //a = reverse(a)
    //b = reverse(b)
    dummy := &ListNode{-1, nil}
    pre := dummy
    add := 0
    for a != nil || b != nil || add > 0 {
        x := add
        if a != nil {
            x += a.Val
            a = a.Next
        }
        if b != nil {
            x += b.Val
            b = b.Next
        }
        add = x / 10
        x = x % 10
        pre.Next = &ListNode{x, nil}
        pre = pre.Next
    }
    return dummy.Next
}
