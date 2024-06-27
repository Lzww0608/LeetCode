func minimumVisitedCells(grid [][]int) int {
    col := make([]hp, len(grid[0]))
    row := hp{}
    f := 1

    for i, r := range grid {
        row = row[:0]
        for j, g := range r {
            for len(row) > 0 && row[0].idx < j {
                heap.Pop(&row)
            }

            c := col[j]
            for len(c) > 0 && c[0].idx < i {
                heap.Pop(&c)
            }

            if i > 0 || j > 0 {
                f = math.MaxInt32
            }

            if len(row) > 0 {
                f = row[0].f + 1
            }

            if len(c) > 0 {
                f = min(f, c[0].f + 1)
            }

            if g > 0 && f < math.MaxInt32 {
                heap.Push(&row, pair{f, g + j})
                heap.Push(&c, pair{f, g + i})
            }

            col[j] = c
        }
    }

    if f == math.MaxInt32 {
        return -1
    }

    return f
}


type pair struct {f, idx int}
type hp []pair
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].f < h[j].f}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(v any) {*h = append(*h, v.(pair))}
func (h *hp) Pop() any { a:= *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v} 
