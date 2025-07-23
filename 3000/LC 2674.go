/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitCircularLinkedList(list *ListNode) []*ListNode {
    p, q := list, list
    for q.Next != list && q.Next.Next != list {
        p = p.Next
        q = q.Next.Next
    }

    if q.Next != list {
        q = q.Next
    }
    q.Next = p.Next
    q = q.Next
    p.Next = list

    return []*ListNode{list, q}
}