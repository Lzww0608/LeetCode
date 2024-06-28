/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBookList(head *ListNode) (ans []int) {

    for head != nil {
        ans = append(ans, head.Val)
        head = head.Next
    }

    for i := 0; i * 2 < len(ans); i++ {
        ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
    }
    
    return
}
