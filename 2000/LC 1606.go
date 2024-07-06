func busiestServers(k int, arrival []int, load []int) (ans []int) {
    used := ListHeap{}
    hp := IntHeap{}

    cnt := make([]int, k)
    for i := 0; i < k; i++ {
        heap.Push(&hp, i)
    }

    for i := range arrival {
        a, b := arrival[i], load[i]
        for used.Len() > 0 && used[0][0] <= a {
            cur := heap.Pop(&used).([]int)
            heap.Push(&hp, i + ((cur[1] - i) % k + k) % k)
        }
        if hp.Len() > 0 {
            idx := heap.Pop(&hp).(int) % k
            cnt[idx]++
            heap.Push(&used, []int{a + b, idx})
        }
    }

    mx := 0
    for i, x := range cnt {
        if x > mx {
            mx = x
            ans = []int{i}
        } else if x == mx {
            ans = append(ans, i)
        }
    }
    
    return 
}

type IntHeap []int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] < h[j]}
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

type ListHeap [][]int
func (h ListHeap) Len() int {return len(h)}
func (h ListHeap) Less(i, j int) bool {return h[i][0] < h[j][0]}
func (h ListHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *ListHeap) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *ListHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

