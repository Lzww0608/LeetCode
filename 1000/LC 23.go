/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    return merge(lists, 0, n - 1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
    if l == r {
        return lists[l]
    }

    mid := l + ((r - l) >> 1)
    a := merge(lists, l, mid)
    b := merge(lists, mid + 1, r)
    return Sort(a, b) 
}

func Sort(a, b *ListNode) *ListNode {
    if a == nil {
        return b 
    } else if b == nil {
        return a
    }
    
    if a.Val < b.Val {
        a.Next = Sort(a.Next, b)
        return a 
    }

    b.Next = Sort(a, b.Next)
    return b
}