
type IntHeap []int
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {
    *h  = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MedianFinder struct {
    pq1 IntHeap
    pq2 IntHeap
}


func Constructor() MedianFinder {
    return MedianFinder{IntHeap{}, IntHeap{}}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.pq1.Len() == 0 || this.pq1[0] >= num {
        heap.Push(&this.pq1, num)
    } else {
        heap.Push(&this.pq2, -num) 
    }
    
    for this.pq2.Len() + 1 < this.pq1.Len() {
        t := -heap.Pop(&this.pq1).(int)
        heap.Push(&this.pq2, t)
    }
    for this.pq1.Len() < this.pq2.Len() {
        t := -heap.Pop(&this.pq2).(int)
        heap.Push(&this.pq1, t)
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.pq1.Len() > this.pq2.Len() {
        return float64(this.pq1[0])
    }
    return float64(this.pq1[0] - this.pq2[0]) / 2.0
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
