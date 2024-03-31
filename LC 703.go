type KthLargest struct {
    sort.IntSlice
    k int
}

func (pq *KthLargest) Push(v interface{}) {
    pq.IntSlice = append(pq.IntSlice, v.(int))
}

func(pq *KthLargest) Pop() interface{} {
    a := pq.IntSlice
    v := a[len(a)-1]
    pq.IntSlice = a[:len(a)-1]
    return v
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