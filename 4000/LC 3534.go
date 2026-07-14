func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
    rk := make([]int, n)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }

    sort.Slice(id, func(i, j int) bool {
        return nums[id[i]] < nums[id[j]]
    })
    for i := range n {
        rk[id[i]] = i
    }

    m := bits.Len(uint(n))
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, m)
    }

    for l, r := 0, 0; r < n; r++ {
        for l < r && nums[id[r]] - nums[id[l]] > maxDiff {
            l++
        }

        f[r][0] = l
    }

    for i := range m - 1 {
        for j := range n {
            p := f[j][i]
            f[j][i + 1] = f[p][i]
        }
    }

    ans := make([]int, len(queries))
    for i, v := range queries {
        l, r := rk[v[0]], rk[v[1]]
        if l == r {
            continue
        }

        if l > r {
            l, r = r, l
        }

        res := 0
        for j := m - 1; j >= 0; j-- {
            if f[r][j] > l {
                res |= 1 << j 
                r = f[r][j]
            }
        }
        //fmt.Println(l, r, res)

        if f[r][0] > l {
            ans[i] = -1
        } else {
            ans[i] = res + 1
        }
    }

    return ans
}