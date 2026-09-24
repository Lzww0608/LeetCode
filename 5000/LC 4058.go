func maxValue(nums []int) int64 {
    n := len(nums)
    sum := 0
    for i, x := range nums {
        if i & 1 == 0 {
            sum += x
        } else {
            sum -= x
        }
    }

    ans := sum 
    f := make([]int, n + 1)
    for i := 1; i < n; i++ {
        d := nums[i] - nums[i - 1]
        if i & 1 == 0 {
            d = -d
        } 

        f[i + 1] = max(f[i - 1] + d * 2, 0)
        ans = max(ans, sum + f[i + 1])
    }

    return int64(ans)
}