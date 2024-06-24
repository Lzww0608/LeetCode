func distributeCookies(cookies []int, k int) int {
    ans, n := int(1e6), len(cookies)
    sort.Ints(cookies)
    bucket := make([]int, k)
    var dfs func(int)
    dfs = func(idx int) {
        if idx == n {
            res := 0
            for _, x := range bucket {
                res = max(res, x)
            }
            ans = min(res, ans)
            return
        }
        for i := 0; i < k; i++ {
            if i > 0 && bucket[i] == bucket[i-1] {
                continue
            }
            if i > 0 && idx == 0 {
                return
            }
            if bucket[i] + cookies[idx] > ans {
                continue
            }
            bucket[i] += cookies[idx]
            dfs(idx + 1)
            bucket[i] -= cookies[idx]
        }
    }
    dfs(0)
    return ans
}
