const MOD int = 1e9 + 7
func getNumberOfBacklogOrders(orders [][]int) (ans int) {
    sell, buy := hp{}, hp{}
    for _, e := range orders {
        p, a, t := e[0], e[1], e[2]
        if t == 0 {
            for a > 0 && sell.Len() > 0 && sell[0].p <= p {
                q := heap.Pop(&sell).(pair)
                x, y := q.p, q.a
                if a >= y {
                    a -= y
                } else {
                    heap.Push(&sell, pair{x, y - a})
                    a = 0
                }
            }
            if a > 0 {
                heap.Push(&buy, pair{-p, a})
            }
        } else {
            for a > 0 && buy.Len() > 0 && -buy[0].p >= p {
                q := heap.Pop(&buy).(pair)
                x, y := q.p, q.a
                if a >= y {
                    a -= y
                } else {
                    heap.Push(&buy, pair{x, y - a})
                    a = 0
                }
            }
            if a > 0 {
                heap.Push(&sell, pair{p, a})
            }
        }
    }

    for sell.Len() > 0 {
        ans = (ans + heap.Pop(&sell).(pair).a) % MOD
    }
    for buy.Len() > 0 {
        ans = (ans + heap.Pop(&buy).(pair).a) % MOD
    }

    return
}

type pair struct{p, a int}

type hp []pair
//maxheap
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool  { return h[i].p < h[j].p }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
