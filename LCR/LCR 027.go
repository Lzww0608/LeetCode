/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

    mid := head
    for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
        mid = mid.Next
    }

    var pre *ListNode
    for mid != nil {
        next := mid.Next
        mid.Next = pre
        pre = mid
        mid = next
    }

    for pre != nil {
        if pre.Val != head.Val {
            return false
        }
        pre = pre.Next
        head = head.Next
    }

    return true
}