import (
    "sort"
)

const maxHeight int = 100

func countRectangles(a [][]int, b [][]int) []int {
    m := len(b)
    h := map[int][]int{}
    for _, v := range a {
        h[v[1]] = append(h[v[1]], v[0])
    }

    id := make([]int, m)
    for i := 0; i < m; i++ {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return b[id[i]][1] > b[id[j]][1]
    })

    
    ans := make([]int, m)
    d := []int{}
    pre := maxHeight
    for _, i := range id {
        x, y := b[i][0], b[i][1]

        f := false
        for t := y; t <= pre; t++ {
            if v, ok := h[t]; ok {
                d = append(d, v...)
                f = true
                delete(h, t)
            }
        }

        pre = y
        
        if f {
            sort.Ints(d)
        }
        
        ans[i] = len(d) - sort.SearchInts(d, x)
    }

    return ans
}