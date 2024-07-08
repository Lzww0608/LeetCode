func minimumDistance(points [][]int) int {
    type pair struct {
        x, y int
    }

    ans := math.MaxInt

    cal := func(idx int) []int {
        a := []pair{}
        b := []pair{}
        for i, v := range points {
            if i != idx {
                a = append(a, pair{v[0] - v[1], i})
                b = append(b, pair{v[0] + v[1], i})
            }
        }
        sort.Slice(a, func(i, j int) bool {
            return a[i].x < a[j].x || (a[i].x == a[j].x && a[i].y < a[j].y)
        })
        sort.Slice(b, func(i, j int) bool {
            return b[i].x < b[j].x || (b[i].x == b[j].x && b[i].y < b[j].y)
        })
        d := max(a[len(a)-1].x - a[0].x, b[len(a)-1].x - b[0].x)
        ans = min(ans, d)
        return []int{a[len(a)-1].y, a[0].y, b[len(b)-1].y, b[0].y}
    }

    for _, x := range cal(-1) {
        cal(x)
    }

    return ans
}