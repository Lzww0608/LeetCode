const N int = 3
func minimumOperationsToWriteY(grid [][]int) int {
    n := len(grid)
    all := [N]int{}
    cnt := [N]int{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i < n / 2 && (i == j || i == n - j - 1) || i >= n / 2 && j == n / 2 {
                cnt[grid[i][j]]++
            } else {
                all[grid[i][j]]++
            }
        }
    }

    ans := n * n
    cntSum := cnt[0] + cnt[1] + cnt[2] 
    for i := 0; i < N; i++ {
        cur := cntSum - cnt[i] + all[i] + min(all[(i + 1) % N], all[(i - 1 + N) % N])
        ans = min(ans, cur)
    }
    
    return ans
}