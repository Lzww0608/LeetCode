/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    arr := []*ListNode{}
    p := head
    for p != nil {
        arr = append(arr, p)
        p = p.Next
    }
    p = head
    i, j := 1, len(arr)-1
    for i <= j {
        if i != j{
            p.Next = arr[j]
            p = p.Next
            j--
        }
        p.Next = arr[i]
        p = p.Next
        i++
    }
    p.Next = nil
    return
}
