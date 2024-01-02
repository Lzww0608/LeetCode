/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    dumb := &ListNode{0, nil}
    arr := []int{}
    
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    arr = append(arr, 0)

    
    i, node := len(arr) - 2, dumb.Next
    for i >= 0 {
        if arr[i+1] <= arr[i] {
            dumb.Next = &ListNode{arr[i], node}
            node = dumb.Next
        } else {
            arr[i] = arr[i+1]
        }
        i--
    }
    return dumb.Next
}