func minimumScore(nums []int, edges [][]int) int {
    n := len(nums)
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    time := 0
    in := make([]int, n)
    out := make([]int, n)
    xor := make([]int, n)
    var dfs func(int, int)
    dfs = func(u, fa int) {
        xor[u] = nums[u]
        
        in[u] = time
        time++
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            dfs(v, u)
            xor[u] ^= xor[v]
        }
        out[u] = time
    }
    dfs(0, -1)

    ans := math.MaxInt32
    check := func(i, j int) bool {
        return in[i] < in[j] && in[j] < out[i]
    }

    var x, y, z int
    for i := 2; i < n; i++ {
        for j := 1; j < i; j++ {
            if check(i, j) {
                x, y, z = xor[i] ^ xor[j], xor[j], xor[0] ^ xor[i]
            } else if check(j, i) {
                x, y, z = xor[i], xor[j] ^ xor[i], xor[0] ^ xor[j]
            } else {
                x, y, z = xor[i], xor[j], xor[0] ^ xor[i] ^ xor[j]
            }

            ans = min(ans, max(x, y, z) - min(x, y, z))
            if ans == 0 {
                return 0
            }
        }
    }

    return ans
}