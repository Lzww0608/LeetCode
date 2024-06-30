func reorganizeString(s string) string {
    n := len(s)
    a := make([]int, 26)
    for i := range s {
        c := int(s[i] - 'a')
        a[c]++
        if a[c] > (n + 1) / 2 {
            return ""
        }
    }
    h := IntHeap{}
    heap.Init(&h)
    for i, x := range a {
        if x > 0 {
            heap.Push(&h, pair{x, i})
        }
    }
    builder := strings.Builder{}
    pre := -1
    for h.Len() > 0 {
        t := heap.Pop(&h).(pair)
        if t.y == pre {
            tt := heap.Pop(&h).(pair)
            pre = tt.y
            builder.WriteByte(byte('a' + tt.y))
            tt.x--
            if tt.x > 0 {
                heap.Push(&h, tt)
            }
            heap.Push(&h, t)
        } else {
            builder.WriteByte(byte('a' + t.y))
            t.x--
            pre = t.y
            if t.x > 0 {
                heap.Push(&h, t)
            }
        }
    }

    return builder.String()
}

type pair struct {
    x, y int
}
type IntHeap []pair
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i].x > h[j].x}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {*h = append(*h, x.(pair))}
func (h * IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
