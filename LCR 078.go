/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := hp{}
    for _, head := range lists {
        if head != nil {
            h = append(h, head)
        }
    }
    heap.Init(&h)
    dummy := &ListNode{-1, nil}
    p := dummy
    for h.Len() > 0 {
        p.Next = h[0]
        p = p.Next
        if h[0].Next != nil {
            h[0] = h[0].Next
            heap.Fix(&h, 0)
        } else {
            heap.Pop(&h)
        }
    }
    return dummy.Next
}

type hp []*ListNode
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].Val < h[j].Val}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(*ListNode))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}