const MOD int = 1_000_000_007
func minMaxSums(nums []int, k int) (ans int) {
    n := len(nums)
    sort.Ints(nums)

    C := make([][]int, n + 1)
    for i := range C {
        C[i] = make([]int, k + 1)
    }
    
    for i := 0; i <= n; i++ {
        C[i][0] = 1
        for j := 1; j < min(k, i + 1); j++ {
            C[i][j] = C[i - 1][j - 1] + C[i - 1][j]
            C[i][j] %= MOD
        }
    }

    sumC := make([]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j < k && j <= i; j++ {
			sumC[i] = (sumC[i] + C[i][j]) % MOD
		}
	}

    for i, x := range nums {
        ans = (ans + x * (sumC[i] + sumC[n - i - 1]) % MOD) % MOD
    }

    return 
}