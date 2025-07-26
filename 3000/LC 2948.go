type pair struct {
    x, i int
}
func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    b := make([]pair, n)
    for i, x := range nums {
        b[i] = pair{x, i}
    }

    sort.Slice(b, func(i, j int) bool {
        return b[i].x < b[j].x
    })
    a := [][]pair{}
    tmp := []pair{}
    for _, v := range b {
        if len(tmp) == 0 || tmp[len(tmp)-1].x + limit >= v.x {
            tmp = append(tmp, v)
        } else {
            a = append(a, append([]pair(nil), tmp...))
            tmp = []pair{v}
        }
    }

    a = append(a, append([]pair(nil), tmp...))

    ans := make([]int, n)
    for _, v := range a {
        m := len(v)
        tmp := make([]int, m)
        for i := 0; i < m; i++ {
            tmp[i] = v[i].i
        }
        sort.Ints(tmp)
        for i := 0; i < m; i++ {
            ans[tmp[i]] = v[i].x
        }
    }

    return ans
}