

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
    n := len(classes)
    h := pairHeap(classes)
    heap.Init(&h)
    for extraStudents > 0 {
        h[0][0]++
        h[0][1]++
        heap.Fix(&h, 0)
        extraStudents--
    }
    for h.Len() > 0 {
        t := heap.Pop(&h).([]int)
        ans += float64(t[0]) / float64(t[1])
    }

    return ans / float64(n)
}



type pairHeap [][]int

func (h pairHeap) Len() int { return len(h) }

func (h pairHeap) Less(i, j int) bool {
    // 计算增加一个学生后的增量比较
    // 增加一个学生后的通过率增量 = (pass+1)/(all+1) - pass/all
    return float64(h[i][0]+1)/float64(h[i][1]+1)-float64(h[i][0])/float64(h[i][1]) >
        float64(h[j][0]+1)/float64(h[j][1]+1)-float64(h[j][0])/float64(h[j][1])
}

func (h pairHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x any) {
    *h = append(*h, x.([]int))
}

func (h *pairHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

