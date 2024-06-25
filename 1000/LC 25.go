/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{-1, head}
    prev, p := dummy, dummy.Next
    n := 0
    for head != nil {
        n++
        head = head.Next
    }
    for i := 0; i < n / k; i++ {
        t := p
        prev.Next = nil
        for j := 0; j < k; j++ {
            next := p.Next
            p.Next = prev.Next
            prev.Next = p 
            p = next
        }
        prev = t
    }
    if prev != nil {
        prev.Next = p
    }

    return dummy.Next
}
