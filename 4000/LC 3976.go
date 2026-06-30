func maxSubarraySum(nums []int, k int) int64 {
    solve := func(f bool) int {
        ans := math.MinInt32
        // a 没有修改
        // b 当前修改过
        // c 当前未修改，之前修改过
        a, b, c := 0, 0, 0
        for _, x := range nums {
            y := x * k 
            if f {
                y = x / k 
            }

            c = max(b, c) + x
            b = max(a, b, 0) + y 
            a = max(a, 0) + x
            ans = max(ans, b, c)
        }

        return ans
    }

    return int64(max(solve(true), solve(false)))
}