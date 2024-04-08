/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    dummy := &ListNode{-1, head}
    preL, preR, l, r := dummy, dummy, dummy.Next, dummy.Next
    for i := 1; i < k; i++ {
        preL, l = preL.Next, l.Next
    }

    for cur := l; cur.Next != nil; cur = cur.Next {
        preR, r = preR.Next, r.Next
    }

    if r.Next == l {
        r.Next, preR.Next = l.Next, l 
        l.Next = r
    } else if l.Next == r {
        l.Next, preL.Next = r.Next, r
        r.Next = l
    } else {
        preL.Next, preR.Next = r, l
        r.Next, l.Next = l.Next, r.Next
    }

    return dummy.Next
}