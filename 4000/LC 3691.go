func maxTotalValue(nums []int, k int) int64 {
    n := len(nums)
    st := NewST(nums)

    h := &hp{}
    for i := range n {
        heap.Push(h, pair{st.query(i, n), i, n})
    }

    ans := 0
    for k > 0 && h.Len() > 0 {
        cur := heap.Pop(h).(pair)
        d, l, r := cur.x, cur.l, cur.r 
        ans += d 
        k--

        if d == 0 {
            break
        }
        heap.Push(h, pair{st.query(l, r - 1), l, r - 1})
    }

    return int64(ans)
}

type ST struct {
    mn, mx [][]int
    n int
}

func NewST(nums []int) ST {
    n := len(nums)
    w := bits.Len(uint(n))
    mn := make([][]int, w)
    mx := make([][]int, w)
    for i := range w {
        mn[i] = make([]int, n)
        mx[i] = make([]int, n)
    }

    for i := range n {
        mn[0][i] = nums[i]
        mx[0][i] = nums[i]
    }

    for i := 1; i < w; i++ {
        for j := 0; j + (1 << (i - 1)) < n; j++ {
            mn[i][j] = min(mn[i - 1][j], mn[i - 1][j + (1 << (i - 1))])
            mx[i][j] = max(mx[i - 1][j], mx[i - 1][j + (1 << (i - 1))])
        }
    }

    return ST{mn, mx, n}
}

func (st *ST)query(l, r int) int {
    w := bits.Len(uint(r - l)) - 1
    mx := max(st.mx[w][l], st.mx[w][r - (1 << w)])
    mn := min(st.mn[w][l], st.mn[w][r - (1 << w)])
    
    return mx - mn
}

type pair struct {
    x, l, r int
}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].x > h[j].x}
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