/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    p, q := head, head.Next
    evenHead := head.Next

    for q != nil && q.Next != nil {
        p.Next = q.Next
        p = p.Next
        q.Next = p.Next
        q = q.Next
    }
    p.Next = evenHead

    return head
}