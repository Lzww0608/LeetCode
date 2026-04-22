func colorGrid(n int, m int, a [][]int) [][]int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    sort.Slice(a, func(i, j int) bool {
        return a[i][2] > a[j][2]
    })

    ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, m)
    }

    for _, v := range a {
        ans[v[0]][v[1]] = v[2]
    }

    for len(a) > 0 {
        cur := a[0]
        a = a[1:]
        i, j := cur[0], cur[1]

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < n && y >= 0 && y < m && ans[x][y] == 0 {
                ans[x][y] = cur[2]
                a = append(a, []int{x, y, cur[2]})
            }
        }
    }

    return ans
}