type KthLargest struct {
    sort.IntSlice
    k int
}

func (pq *KthLargest) Push(x any) {
    pq.IntSlice = append(pq.IntSlice, x.(int))
}

func (pq *KthLargest) Pop() (x any) {
    old := pq.IntSlice
    n := len(old)
    x = old[n-1]
    pq.IntSlice = old[:n-1]
    return
}


func Constructor(k int, nums []int) KthLargest {
    pq := KthLargest{k: k}
    for _, x := range nums {
        pq.Add(x)
    }
    return pq
}


func (pq *KthLargest) Add(val int) int {
    heap.Push(pq, val)
    if pq.Len() > pq.k {
        heap.Pop(pq)
    }
    return pq.IntSlice[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */