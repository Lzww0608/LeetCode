func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) (ans []int) {
    m := map[string]int{}
    for _, s := range positive_feedback {
        m[s] = 3
    }
    for _, s := range negative_feedback {
        m[s] = -1
    }
    h := hp{}
    heap.Init(&h)
    for i, v := range report {
        score, id := 0, student_id[i]
        str := strings.Split(v, " ")
        for _, s := range str {
            score += m[s]
        }
        heap.Push(&h, []int{id, score})
        if h.Len() > k {
            heap.Pop(&h)
        }
    }
    for h.Len() > 0 {
        ans = append(ans, heap.Pop(&h).([]int)[0])
    }
    for i, j := 0, k - 1; i < j; i, j = i + 1, j - 1 {
        ans[i], ans[j] = ans[j], ans[i]
    }

    return
}

type hp [][]int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i][1] < h[j][1] || (h[i][1] == h[j][1] && h[i][0] > h[j][0])}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.([]int))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}