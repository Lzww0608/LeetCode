
import (
	"container/heap"
)
const MOD int = 1_000_000_007
type pair struct {
    x, i int
}
func getFinalState(nums []int, k int, multiplier int) []int {
    if multiplier == 1 {
        return nums
    }
    mx := 0
    n := len(nums)
    h := make(hp, n)
    for i, x := range nums {
        mx = max(mx, x)
        h[i] = pair{x, i}
    }
    heap.Init(&h)

    for ; k > 0 && h[0].x < mx; k-- {
        h[0].x *= multiplier
        heap.Fix(&h, 0)
    }

    sort.Slice(h, func(i, j int) bool {
        return h[i].x < h[j].x || h[i].x == h[j].x && h[i].i < h[j].i
    })
    for i, p := range h {
        e := k / n 
        if i < k % n {
            e++
        }
        nums[p.i] = p.x % MOD * quickPow(multiplier, e) % MOD
    }

    return nums
}


type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].x < h[j].x || h[i].x == h[j].x && h[i].i < h[j].i}
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

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        a = (a * a) % MOD
        r >>= 1
    }
    return res
}
