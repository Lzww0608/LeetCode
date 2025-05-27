/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }

    dummy := &ListNode{-1, head}
    slow, fast := dummy, dummy
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    p := head
    q := reverse(slow.Next)
    
    for q != nil {
        next := p.Next
        p.Next = q
        q = q.Next
        p.Next.Next = next
        p = next
    }

    if p != nil {
        p.Next = nil
    }
    return 
}

func reverse(head *ListNode) *ListNode {
    var pre *ListNode
    p := head 
    for p != nil {
        next := p.Next
        p.Next = pre 
        pre = p 
        p = next
    }

    return pre
}