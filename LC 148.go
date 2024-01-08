/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//归并排序：自顶向下
func mergeList(p, q *ListNode) *ListNode {
    dummy := &ListNode{0, nil}
    t := dummy
    for p != nil && q != nil {
        if p.Val <= q.Val {
            t.Next = &ListNode{p.Val, nil}
            p = p.Next
        } else {
            t.Next = &ListNode{q.Val, nil}
            q = q.Next
        }
        t = t.Next
    }
    if p != nil {
        t.Next = p
    } else {
        t.Next = q
    }
    return dummy.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    fast, slow.Next, slow = slow.Next, nil, head
    slow, fast = sortList(slow), sortList(fast)
    return mergeList(slow, fast)
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//归并排序：自底向上
func mergeList(p, q *ListNode) *ListNode {
    dummy := &ListNode{0, nil}
    t := dummy
    for p != nil && q != nil {
        if p.Val <= q.Val {
            t.Next = &ListNode{p.Val, nil}
            p = p.Next
        } else {
            t.Next = &ListNode{q.Val, nil}
            q = q.Next
        }
        t = t.Next
    }
    if p != nil {
        t.Next = p
    } else {
        t.Next = q
    }
    return dummy.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p, n := head, 0
    for p != nil {
        n++
        p = p.Next
    }
    dummy := &ListNode{0, head}
    for curLen := 1; curLen < n; curLen <<= 1 {
        cur, prev, cnt := dummy.Next, dummy, 0
        for cur != nil && cnt + curLen <= n  {
            l := cur
            var r, p *ListNode
            for i := 1; i <= curLen && cur != nil; i++ {
                p = cur
                cur = cur.Next
            }
            p.Next, r = nil, cur
            for i := 1; i <= curLen && cur != nil; i++ {
                p = cur
                cur = cur.Next
            }
            p.Next = nil
            prev.Next = mergeList(l,r)
            for prev.Next != nil {
                prev = prev.Next
            }
            prev.Next = cur
        }
    }
    return dummy.Next
}