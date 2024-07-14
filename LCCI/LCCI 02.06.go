/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    n := 0
    for p := head; p != nil; p = p.Next {
        n++
    }

    p := head
    for i := 1; i < (n + 1) / 2; i++ {
        p = p.Next
    }

    q := p.Next
    p.Next = nil
    for q != nil {
        next := q.Next
        q.Next = p.Next
        p.Next = q
        q = next
    }

    for p = p.Next; p != nil; p, head = p.Next, head.Next {
        if head.Val != p.Val {
            return false
        }
    }

    return true
}

