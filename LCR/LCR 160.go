import (
    "container/heap"
)

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

type MedianFinder struct {
    pq1 hp
    pq2 hp
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    a, b := hp{}, hp{}
    heap.Init(&a)
    heap.Init(&b)
    return MedianFinder{a, b}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(&this.pq1, num)
    heap.Push(&this.pq2, -heap.Pop(&this.pq1).(int))
    if this.pq1.Len() < this.pq2.Len() {
        heap.Push(&this.pq1, -heap.Pop(&this.pq2).(int))
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
