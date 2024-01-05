/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    slow, fast := head, head
    for fast != nil && fast.Next != nil && slow != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            slow = head
            for fast != slow {
                slow = slow.Next
                fast = fast.Next
            }
            return fast
        }
    }
    return nil
}