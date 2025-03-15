func maximumProcessableQueries(nums []int, queries []int) (ans int) {
    n, m := len(nums), len(queries)
    f := make([][]int, n + 2)
    for i := range f {
        f[i] = make([]int, n + 2)
    }
    
    for l := 0; l < n; l++ {
        for r := n - 1; r >= l; r-- {
            cur := 0
            if l > 0 {
                if nums[l-1] >= queries[f[l-1][r]] {
                    cur = max(cur, f[l-1][r] + 1)
                } else {
                    cur = f[l-1][r]
                }
            }
            
            if r < n - 1 {
                if nums[r+1] >= queries[f[l][r+1]] {
                    cur = max(cur, f[l][r+1] + 1)
                } else {
                    cur = max(cur, f[l][r+1])
                }
            }
            
            if cur == m || cur == n {
                return cur
            }
            f[l][r] = cur
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] >= queries[f[i][i]] {
            ans = max(ans, f[i][i] + 1)
        } else {
            ans = max(ans, f[i][i])
        }
    }

    return
}