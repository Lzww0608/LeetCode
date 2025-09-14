func modeWeight(nums []int, k int) int64 {
    ans := 0
    cnt := make(map[int]int)
    h := &hp{}
    for l, r := 0, 0; r < len(nums); r++ {
        x := nums[r]
        cnt[x]++
        heap.Push(h, pair{x, cnt[x]})

        if r > k - 1 {
            x = nums[l]
            cnt[x]--
            heap.Push(h, pair{x, cnt[x]})
            l++
        }

        for cnt[(*h)[0].x] != (*h)[0].cnt {
            heap.Pop(h)
        } 

        if r >= k - 1 {
            ans += (*h)[0].x * (*h)[0].cnt 
        }
        
    }

    return int64(ans)
}

type pair struct {
    x, cnt int 
}

type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {
    if h[i].cnt == h[j].cnt {
        return h[i].x < h[j].x
    }

    return h[i].cnt > h[j].cnt
}
func (h hp) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
} 
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[: n - 1]
    return
}