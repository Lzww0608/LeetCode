/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortLinkedList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    } 

    pre, cur := head, head.Next
    for cur != nil {
        for cur != nil && cur.Val >= 0 {
            cur = cur.Next
            pre = pre.Next
        }

        if cur != nil {
            pre.Next = cur.Next
            cur.Next = head 
            head = cur 
            cur = pre.Next
        }
    }

    return head
}

