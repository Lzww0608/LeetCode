
import "container/heap"
func findCrossingTime(n int, k int, time [][]int) (cur int) {
    sort.SliceStable(time, func(i, j int) bool {
        a, b := time[i], time[j]
        return a[0] + a[2] < b[0] + b[2]
    })

    waitL, waitR := make(hp, k), hp{}
    for i := range waitL {
        waitL[i].i = k - 1 - i
    }
    workL, workR := hp2{}, hp2{}

    for n > 0 {
        for len(workL) > 0 && workL[0].t <= cur {
            heap.Push(&waitL, heap.Pop(&workL))
        }
        for len(workR) > 0 && workR[0].t <= cur {
            heap.Push(&waitR, heap.Pop(&workR))
        }

        if len(waitR) > 0 {
            p := heap.Pop(&waitR).(pair)
            cur += time[p.i][2]
            heap.Push(&workL, pair{p.i, cur + time[p.i][3]})
        } else if len(waitL) > 0 {
            p := heap.Pop(&waitL).(pair)
            cur += time[p.i][0]
            heap.Push(&workR, pair{p.i, cur + time[p.i][1]})
            n--
        } else if len(workL) == 0 {
            cur = workR[0].t
        } else if len(workR) == 0 {
            cur = workL[0].t;
        } else {
            cur = min(workL[0].t, workR[0].t)
        }
    }

    for len(workR) > 0 {
        p := heap.Pop(&workR).(pair)
        cur = max(p.t, cur) + time[p.i][2]
    }

    return 
}

type pair struct{i, t int}

type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].i > h[j].i}
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

type hp2 []pair
func (h hp2) Len() int {return len(h)}
func (h hp2) Less(i, j int) bool {return h[i].t < h[j].t}
func (h hp2) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp2) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp2) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}