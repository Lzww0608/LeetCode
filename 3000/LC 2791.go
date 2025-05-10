func countPalindromePaths(parent []int, s string) int64 {
    n := len(parent)
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        x := parent[i]
        g[x] = append(g[x], i)
    }

    ans := 0
    cnt := make(map[int]int)
    cnt[0] = 1
    var dfs func(int, int)
    dfs = func(u, xor int) {
        for _, v := range g[u] {
            t := xor ^ (1 << int(s[v] - 'a'))
            ans += cnt[t]
            for i := 0; i < 26; i++ {
                ans += cnt[t ^ (1 << i)]
            }
            cnt[t]++
            dfs(v, t)
        }
    }

    dfs(0, 0)

    return int64(ans)
}