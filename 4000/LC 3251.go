
const MOD int = 1_000_000_007
const M int = 1005
func countOfPairs(nums []int) (ans int) {
    n := len(nums)

    f := make([][M]int, n)
    s := make([]int, M)

    for j := range f[0] {
        f[0][j] = 1
    }

    for i := 1; i < n; i++ {
        s[0] = f[i-1][0]
        for k := 1; k < M; k++ {
            s[k] = s[k-1] + f[i-1][k]
        }

        for j := 0; j <= nums[i]; j++ {
            mx := j + min(nums[i-1] - nums[i], 0)
            if mx >= 0 {
                f[i][j] = s[mx] % MOD
            }
        }
    }

    for _, x := range f[n-1][:nums[n-1]+1] {
        ans = (ans + x) % MOD
    }

    return 
}