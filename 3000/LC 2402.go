import "container/heap"
import "sort"
func mostBooked(n int, meetings [][]int) (ans int) {
    cnt := make([]int, n)
    h := &hp{}
    ava := &hp{}
    heap.Init(ava)
    heap.Init(h)
    for i := 0; i < n; i++ {
        heap.Push(ava, pair{i, 0})
    }

    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][0] < meetings[j][0]
    })

    for _, v := range meetings {
        start, end := v[0], v[1]
        for h.Len() > 0 && (*h)[0].end <= start {
            tmp := heap.Pop(h).(pair)
            tmp.end = 0
            heap.Push(ava, tmp)
        }
        var cur pair
        if ava.Len() > 0 {
            cur = heap.Pop(ava).(pair)
        } else {
            cur = heap.Pop(h).(pair)
        } 
        cnt[cur.idx]++
        cur.end = max(cur.end, start) + end - start
        heap.Push(h, cur)
    }
    for i, x := range cnt {
        if x > cnt[ans] {
            ans = i
        }
    }

    return
}

type pair struct {
    idx, end int
}

type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].end < h[j].end || h[i].end == h[j].end && h[i].idx < h[j].idx}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
} 
func (h *hp) Pop() (x any) {
    old := *h
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}