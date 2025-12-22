func maximumScore(nums []int, s string) int64 {
    ans := 0 
    n := len(nums)

    h := &hp{}
    b := false
    for i := 0; i < n; i++ {
        if s[i] == '0' {
            heap.Push(h, nums[i])
            b = true
        } else {
            if b {
                heap.Push(h, nums[i])
                ans += heap.Pop(h).(int)
            } else {
                ans += nums[i]
            }
        }
    }


    return int64(ans)
}

type hp []int 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() (x any) {
    old := *h  
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}