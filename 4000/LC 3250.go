const MOD int = 1_000_000_007
const M int = 51
func countOfPairs(nums []int) int {
    n := len(nums)
    f := make([][M]int, n + 1)
    for i := range f[n] {
        f[n][i] = 1
    }
    for i := n - 1; i >= 0; i-- {
        for pre := 0; pre < M; pre++ {
            for a := pre; a <= nums[i]; a++ {
                b := nums[i] - a
                if i == 0 || b <= nums[i-1] - pre {
                    f[i][pre] += f[i+1][a]
                    f[i][pre] %= MOD
                } 
            }
        }
    }


    return f[0][0]
}