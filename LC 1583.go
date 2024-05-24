func unhappyFriends(n int, preferences [][]int, pairs [][]int) (ans int) {
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
    }
    for i, v := range preferences {
        for j, x := range v {
            f[i][x] = j
        }
        f[i][i] = 0x3f3f3f3f
    }

    m := map[int]int{}
    for _, v := range pairs {
        m[v[0]] = v[1]
        m[v[1]] = v[0]
    }

    for x := 0; x < n; x++ {
        y := m[x]
        for u := range f[x] {
            if f[x][u] < f[x][y] && f[u][x] < f[u][m[u]] {
                ans++
                break
            }
        }
    }

    return
}