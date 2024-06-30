var mod int = int(1e9 + 7)
func numsGame(nums []int) []int {
    s1, s2 := 0, 0
    h1 := hp{}
    h2 := hp{}
    heap.Init(&h1)
    heap.Init(&h2)
    ans := make([]int, len(nums))
    for i := range nums {
        nums[i] -= i
        x := nums[i]
        if h1.Len() == 0 || h1[0] >= x {
            heap.Push(&h1, x)
            s1 += x
            if h1.Len() > h2.Len() + 1 {
                t := heap.Pop(&h1).(int)
                heap.Push(&h2, -t)
                s1 -= t
                s2 += t
            }
        } else {
            heap.Push(&h2, -x)
            s2 += x
            if h1.Len() < h2.Len() {
                t := -heap.Pop(&h2).(int)
                heap.Push(&h1, t)
                s1 += t
                s2 -= t
            }
        }

        if i & 1 == 1 {
            ans[i] = abs(s2 - s1) % mod
        } else {
            ans[i] = abs(s2 - s1 + h1[0]) % mod
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
} 
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
