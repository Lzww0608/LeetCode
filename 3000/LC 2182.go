func repeatLimitedString(s string, k int) string {
    alp := make([]int, 26)
    for _, c := range s {
        x := int(c - 'a')
        alp[x]++
    }
    h := hp{}
    heap.Init(&h)
    for i, x := range alp {
        if x != 0 {
            heap.Push(&h, pair{i, x})
        }
    }

    ans := []byte{}
    for h.Len() > 0 {
        tmp := heap.Pop(&h).(pair)
        mn := min(k, tmp.y)
        for i := 0; i < mn; i++ {
            ans = append(ans, byte('a' + tmp.x))
        }
        tmp.y -= mn
        if tmp.y > 0 {
            if h.Len() == 0 {
                break
            }
            ans = append(ans, byte('a' + h[0].x))
            h[0].y--
            if h[0].y == 0 {
                heap.Pop(&h)
            }
            heap.Push(&h, tmp)
        } 
    }

    return string(ans)
}

type pair struct {
    x, y int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].x > h[j].x}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    old = old[:n-1]
    *h = old
    return x
}