const N int = 100_000
func minOperations(n int, m int) int {
    isPrime := make([]int, N)
    isPrime[1] = 1
    for i := 2; i < N; i++ {
        if isPrime[i] == 1 {
            continue
        }
        for j := i * i; j < N; j += i {
            isPrime[j] = 1
        }
    }

    if isPrime[n] == 0 || isPrime[m] == 0 {
        return -1
    }

    dis := make([]int, N)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[n] = n 
    h := &hp{{n, n}}
    for h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        u, w := cur.to, cur.weight
        if u == m {
            return w
        } else if w > dis[u] {
            continue
        }
        p := 1
        for v := u; v > 0; v /= 10 {
            d := v % 10 
            if d > 0 {
                x := u - p 
                cur_w := w + x 
                if cur_w < dis[x] && isPrime[x] == 1 {
                    dis[x] = cur_w 
                    heap.Push(h, pair{x, cur_w})
                }
            }
            if d < 9 {
                x := u + p 
                cur_w := w + x 
                if cur_w < dis[x] && isPrime[x] == 1 {
                    dis[x] = cur_w 
                    heap.Push(h, pair{x, cur_w})
                }
            }
            p *= 10
        }
    }

    return -1
}

type pair struct {
    to, weight int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].weight < h[j].weight}
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