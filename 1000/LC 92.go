/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right {
        return head
    }

    dummy := &ListNode{-1, head}
    p, q:= dummy, dummy.Next
    tail := p

    for i := 1; i <= right; i++ {
        next := q.Next
        if i == left {
            tail = q
        }

        if i >= left {
            q.Next = p.Next
            p.Next = q
        } else {
            p = p.Next
        }
        if i == left {
            tail = q
        }
        q = next
    }
    tail.Next = q

    return dummy.Next
}
