import (
    "container/heap"
)

func isPossible(target []int) bool {
    if len(target) == 1 {
        return target[0] == 1
    }

    maxHeap := &IntHeap{}
    heap.Init(maxHeap)
    sum := 0
    for _, val := range target {
        heap.Push(maxHeap, val)
        sum += val
    }

    for {
        maxVal := heap.Pop(maxHeap).(int)
        sum -= maxVal

        if maxVal == 1 || sum == 1 {
            return true
        }

        if maxVal < sum || maxVal%sum == 0 {
            return false
        }

        maxVal %= sum
        sum += maxVal
        heap.Push(maxHeap, maxVal)
    }
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}