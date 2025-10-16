const MOD int = 1_000_000_007
func countPermutations(complexity []int) int {
    n := len(complexity)
    for i := 1; i < n; i++ {
        if complexity[i] <= complexity[0] {
            return 0
        }
    }

    ans := 1
    for i := 1; i < n; i++ {
        ans = ans * i % MOD
    }

    return ans
}