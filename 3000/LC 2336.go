
import "container/heap"
type SmallestInfiniteSet struct {
    m map[int]bool
    h hp
}

const N int = 1_000

func Constructor() SmallestInfiniteSet {
    m := make(map[int]bool)
    h := hp{}
    for i := 1; i <= N; i++ {
        m[i] = true 
        heap.Push(&h, i)
    }
    return SmallestInfiniteSet{m, h}
}


func (c *SmallestInfiniteSet) PopSmallest() int {
    if c.h.Len() == 0 {
        return -1
    }
    x := heap.Pop(&c.h).(int)
    c.m[x] = false
    return x
}


func (c *SmallestInfiniteSet) AddBack(num int)  {
    if !c.m[num] {
        c.m[num] = true 
        heap.Push(&c.h, num)
    }
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] < h[j]}
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
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */