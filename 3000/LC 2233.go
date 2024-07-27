const MOD int = 1_000_000_007

func maximumProduct(nums []int, k int) int {
    h := hp(nums)
    heap.Init(&h)

    ans := 1
    for i := 0; i < k; i++ {
        h[0]++
        heap.Fix(&h, 0)
    }

    for i := 0; i < h.Len(); i++ {
        ans = (ans * h[i]) % MOD
    }

    return ans
}

type hp []int
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
