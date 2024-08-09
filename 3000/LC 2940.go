func leftmostBuildingQueries(heights []int, queries [][]int) []int {
    n := len(queries)
    ans := make([]int, n)
    type pair struct {
        x, i int
    }
    qs := make([][]pair, len(heights))
    for i, v := range queries {
        a, b := v[0], v[1]
        if a > b {
            a, b = b, a
        }
        if a == b || heights[a] < heights[b] {
            ans[i] = b
        } else {
            qs[b] = append(qs[b], pair{heights[a], i})
            ans[i] = -1
        }
    }

    st := []int{}
    for i := len(heights) - 1; i >= 0; i-- {
        for _, v := range qs[i] {
            pos := sort.Search(len(st), func(j int) bool {
                return heights[st[j]] <= v.x
            })
            if pos > 0 {
                ans[v.i] = st[pos-1]
            }
        }

        for len(st) > 0 && heights[i] >= heights[st[len(st)-1]] {
            st = st[:len(st)-1]
        }

        st = append(st, i)
    }

    return ans
}