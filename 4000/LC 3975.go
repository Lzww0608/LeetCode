func filterOccupiedIntervals(a [][]int, freeStart int, freeEnd int) [][]int {
    sort.Slice(a, func(i, j int) bool {
        if a[i][0] == a[j][0] {
            return a[i][1] < a[j][1]
        } 
        return a[i][0] < a[j][0]
    })
    n := len(a)
    b := make([][]int, 0, n)
    for _, v := range a {
        if len(b) == 0 || v[0] > b[len(b) - 1][1] + 1 {
            b = append(b, v)
        } else {
            b[len(b) - 1][1] = max(b[len(b) - 1][1], v[1])
        }
    }

    n = len(b)
    ans := make([][]int, 0, n)
    for _, v := range b {
        if v[0] >= freeStart && v[1] <= freeEnd {
            continue
        }

        if v[1] < freeStart || v[0] > freeEnd {
            ans = append(ans, v)
        } else if v[0] < freeStart {
            ans = append(ans, []int{v[0], freeStart - 1})
            if v[1] > freeEnd {
                ans = append(ans, []int{freeEnd + 1, v[1]})
            }
        } else {
            ans = append(ans, []int{freeEnd + 1, v[1]})
        }
    }

    return ans
}