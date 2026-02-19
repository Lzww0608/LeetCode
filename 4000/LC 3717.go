func minOperations(nums []int) int {
    n := len(nums)
    f := make([][101]int, n)
    for i := range f {
        for j := range f[i] {
            f[i][j] = math.MaxInt32
        }
    }
    f[0][nums[0]] = 0

    for i := 1; i < n; i++ {
        x := nums[i]
        for j := x; j <= 100; j++ {
            for k := 1; k <= 100; k++ {
                if f[i - 1][k] >= math.MaxInt32 {
                    continue
                }
                if j % k == 0 {
                    f[i][j] = min(f[i][j], f[i - 1][k] + j - x)
                }
            }
        }
    }
    
    ans := f[n - 1][1]
    for j := 1; j <= 100; j++ {
        ans = min(ans, f[n - 1][j])
    }
    //fmt.Println(f[n - 1])
    return ans
}