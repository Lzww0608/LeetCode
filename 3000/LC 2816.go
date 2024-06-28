/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head.Val > 4 {
        head = &ListNode{0, head}
    }

    for cur := head; cur != nil; cur = cur.Next {
        cur.Val = cur.Val * 2 % 10
        if cur.Next != nil && cur.Next.Val > 4 {
            cur.Val++
        }
    }

    return head
}
