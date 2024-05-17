func medianSlidingWindow(nums []int, k int) []float64 {
    pq1, pq2 := IntHeap{}, IntHeap{}
    heap.Init(&pq1)
    heap.Init(&pq2)
    m := map[int]int{}
    
    cal := func(k int) float64 {
        if k & 1 == 1 {
            return float64(pq1[0])
        }
        return float64(pq1[0] - pq2[0]) / 2.0
    }

    for i := 0; i < k; i++ {
        heap.Push(&pq1, nums[i])
    }
    for i := 0; i < k / 2; i++ {
        heap.Push(&pq2, -heap.Pop(&pq1).(int))
    }

    ans := make([]float64, 0, len(nums))
    ans = append(ans, cal(k))
    for i := k; i < len(nums); i++ {
        balance := 0
        x, y := nums[i-k], nums[i]
        m[x]++
        if pq1.Len() > 0 && x <= pq1[0] {
            balance--
        } else {
            balance++
        }
        if pq1.Len() > 0 && y <= pq1[0] {
            heap.Push(&pq1, y)
            balance++
        } else {
            heap.Push(&pq2, -y)
            balance--
        }

        if balance > 0 {
            heap.Push(&pq2, -heap.Pop(&pq1).(int))
        } else if balance < 0 {
            heap.Push(&pq1, -heap.Pop(&pq2).(int))
        }

        for pq1.Len() > 0 && m[pq1[0]] > 0 {
            m[pq1[0]]--
            heap.Pop(&pq1)
        }

        for pq2.Len() > 0 && m[-pq2[0]] > 0 {
            m[-pq2[0]]--
            heap.Pop(&pq2)
        }

        ans = append(ans, cal(k))
    }

    return ans
}

type IntHeap []int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}