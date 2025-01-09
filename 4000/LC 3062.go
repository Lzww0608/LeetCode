/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gameResult(head *ListNode) string {
    score := 0
    for head != nil {
        x, y := head.Val, head.Next.Val
        if x < y {
            score++
        } else {
            score--
        }
        head = head.Next.Next
    }

    if score > 0 {
        return "Odd"
    } else if score < 0 {
        return "Even"
    }
    return "Tie"
}