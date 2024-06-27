type MedianFinder struct {
    maxHeap, minHeap hp
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    maxHeap := hp{}
    minHeap := hp{}
    heap.Init(&minHeap)
    heap.Init(&maxHeap)
    return MedianFinder{maxHeap, minHeap}
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxHeap.Len() == 0 || this.maxHeap[0] >= num {
        heap.Push(&this.maxHeap, num)
        if this.maxHeap.Len() > this.minHeap.Len() + 1 {
            heap.Push(&this.minHeap, -heap.Pop(&this.maxHeap).(int))
        }
    } else {
        heap.Push(&this.minHeap, -num)
        if this.maxHeap.Len() < this.minHeap.Len() {
            heap.Push(&this.maxHeap, -heap.Pop(&this.minHeap).(int))
        }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Len() == this.minHeap.Len() {
        a := float64(this.maxHeap[0])
        b := -float64(this.minHeap[0])
        return (a + b) / 2
    } 
    return float64(this.maxHeap[0])
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
    x = old[n-1]
    *h = old[:n-1]
    return 
}



/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */