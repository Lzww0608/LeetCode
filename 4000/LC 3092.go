
import "container/heap"
func mostFrequentIDs(nums []int, freq []int) []int64 {
    n := len(nums)
    ans := make([]int64, n)
    h := &hp{}
    cnt := make(map[int]int)
    for i, x := range nums {
        cnt[x] += freq[i]
        heap.Push(h, [2]int{cnt[x], x})
        for {
            cur := (*h)[0]
            if cur[0] == cnt[cur[1]] {
                ans[i] = int64(cur[0])
                break
            } 
            heap.Pop(h)
        }
    }

    return ans
}


type hp [][2]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][0] > h[j][0]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([2]int))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return
}