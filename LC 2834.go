var MOD int = int(1e9 + 7)
func minimumPossibleSum(n int, target int) int {
    if n <= target >> 1 {
        return n * (n + 1) % MOD / 2 
    }
    
    ans := (target / 2) * (target / 2 + 1) / 2 
    n -= target >> 1
    ans += (2 * target + n - 1) * n / 2
    return ans % MOD
}