const MOD int = 1e9 + 7
func countStableSubsequences(nums []int) (ans int) {
    n := len(nums)
    if n <= 2 {
        return n * (n + 1) / 2
    } 
    
    f := make([][3][3]int, n)
    for i := range f {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                f[i][j][k] = -1
            }
        }
    }
    var dfs func(int, int, int) int 
    dfs = func(i, pre, ppre int) (res int) {
        if i == n {
            return 1
        }

        p := &f[i][pre + 1][ppre + 1]
        if *p != -1 {
            return *p
        }

        if pre == -1 || ppre == -1 {
            res = dfs(i + 1, pre, ppre) % MOD
            res = (res + dfs(i + 1, nums[i] & 1, pre)) % MOD
            *p = res
            return 
        }

        res = dfs(i + 1, pre, ppre) % MOD
        if ppre == pre {
            if pre != nums[i] & 1 {
                res = (res + dfs(i + 1, nums[i] & 1, pre)) % MOD 
            } 
        } else {
            res = (res + dfs(i + 1, nums[i] & 1, pre)) % MOD
        }

        *p = res 
        return
    }
    
    return (dfs(0, -1, -1) + MOD - 1) % MOD
}