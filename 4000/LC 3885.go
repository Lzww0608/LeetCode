type EventManager struct {
    m map[int]int
    h hp
}


func Constructor(events [][]int) EventManager {
    m := make(map[int]int)
    h := hp{}
    for _, v := range events {
        m[v[0]] = v[1]
        heap.Push(&h, pair{v[0], v[1]})
    }

    return EventManager{m, h}
}


func (c *EventManager) UpdatePriority(eventId int, newPriority int)  {
    c.m[eventId] = newPriority
    heap.Push(&c.h, pair{eventId, newPriority})
}


func (c *EventManager) PollHighest() int {
    for c.h.Len() > 0 {
        cur := heap.Pop(&c.h).(pair)
        if c.m[cur.i] == cur.p {
            delete(c.m, cur.i)
            return cur.i
        }
    }

    return -1
}

type pair struct {
    i, p int
}

type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].p > h[j].p || h[i].p == h[j].p && h[i].i < h[j].i}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return 
}


/**
 * Your EventManager object will be instantiated and called as such:
 * obj := Constructor(events);
 * obj.UpdatePriority(eventId,newPriority);
 * param_2 := obj.PollHighest();
 */