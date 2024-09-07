/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    var last *ListNode
    p := head

    for p != nil {
        if p.Val != 9 {
            last = p
        }
        p = p.Next
    }
    if last == nil {
        tmp := &ListNode{1, head}
        head = tmp
        last = head
    } else {
        last.Val += 1
    }

    
    for last = last.Next; last != nil; last = last.Next {
        last.Val = 0
    }

    return head
}