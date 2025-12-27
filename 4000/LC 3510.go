func minimumPairRemoval(nums []int) (ans int) {
    n := len(nums)
    h := hp{}
    left := make([]int, n + 1)
    right := make([]int, n)
    dec := 0
    for i := range n {
        left[i] = i - 1
        right[i] = i + 1
        if i + 1 < n {
            heap.Push(&h, pair{nums[i] + nums[i + 1], i})
            if nums[i] > nums[i + 1] {
                dec++
            }
        }
    }
    remove := func(i int) {
        l, r := left[i], right[i]
        right[l] = r 
        left[r] = l 
    }
    lazy := make(map[pair]int)

    for dec > 0 {
        ans++
        for lazy[h[0]] > 0 {
            lazy[h[0]]--
            heap.Pop(&h)
        }

        p := heap.Pop(&h).(pair)
        s, i := p.s, p.i 
        j := right[i]
        if nums[i] > nums[j] {
            dec--
        }

        pre := left[i]
        if pre >= 0 {
            if nums[pre] > nums[i] {
                dec--
            }
            if nums[pre] > s {
                dec++
            }
            lazy[pair{nums[pre] + nums[i], pre}]++
            heap.Push(&h, pair{nums[pre] + s, pre})
        }

        nxt := right[j]
        if nxt < n {
            if nums[j] > nums[nxt] {
                dec--
            }
            if s > nums[nxt] {
                dec++
            }
            lazy[pair{nums[j] + nums[nxt], j}]++
            heap.Push(&h, pair{nums[nxt] + s, i})
        }

        nums[i] = s 
        remove(j)
    }

    return
}

type pair struct {
    s, i int 
}
type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].s < h[j].s || h[i].s == h[j].s && h[i].i < h[j].i}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(*h)
    x = old[n - 1]
    *h = old[:n - 1]
    return 
}