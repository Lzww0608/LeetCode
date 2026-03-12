func minCost(s string, encCost int, flatCost int) int64 {
    n := len(s)
    pre := make([]int, n + 1)
    for i := range s {
        if s[i] == '1' {
            pre[i + 1] = pre[i] + 1
        } else {
            pre[i + 1] = pre[i]
        }
    }

    var dfs func(int, int) int 
    dfs = func(l, r int) int {
        d := r - l + 1
        cnt := pre[r + 1] - pre[l]
        if cnt == 0 {
            return flatCost
        }
        
        res := d * encCost * cnt
        if d & 1 == 1 {
            return res
        }
        
        mid := l + ((r - l) >> 1)
        res = min(res, dfs(l, mid) + dfs(mid + 1, r))
        return res
    }

    return int64(dfs(0, n - 1))
}