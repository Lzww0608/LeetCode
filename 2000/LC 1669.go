/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    start, end := list1, list1
    for p, i := list1, 0; i <= b; p = p.Next {
        if i + 1 == a {
            start = p 
        }
        if i == b {
            end = p
        }
        i++
    }

    start.Next = list2
    for start.Next != nil {
        start = start.Next
    }
    start.Next = end.Next
    return list1
}