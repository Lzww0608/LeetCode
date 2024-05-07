/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) (ans int) {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
    }
    p, q := slow, slow.Next
    l, r := head, slow.Next
    slow.Next = nil
    for q != nil {
        next := q.Next
        q.Next = p.Next
        p.Next = q
        q = next
    }
    r.Next = nil
    r = slow.Next
    for r != nil {
        ans = max(ans, l.Val + r.Val)
        l, r = l.Next, r.Next
    }

    return
}