func maxSumTrionic(nums []int) int64 {
    n := len(nums)
    ans, f1, f2, f3 := math.MinInt / 2, math.MinInt / 2, math.MinInt / 2, math.MinInt / 2

    for i := 1; i < n; i++ {
        x, y := nums[i], nums[i - 1]
        if x > y {
            f3 = max(f2, f3) + x
        } else {
            f3 = math.MinInt / 2
        }func maxSumTrionic(nums []int) int64 {
    n := len(nums)
    ans, f1, f2, f3 := math.MinInt / 2, math.MinInt / 2, math.MinInt / 2, math.MinInt / 2

    for i := 1; i < n; i++ {
        x, y := nums[i], nums[i - 1]
        if x > y {
            f3 = max(f2, f3) + x
        } else {
            f3 = math.MinInt / 2
        }

        if x < y {
            f2 = max(f2, f1) + x
        } else {
            f2 = math.MinInt / 2
        }

        if x > y {
            f1 = max(f1, y) + x
        } else {
            f1 = math.MinInt / 2
        }
        
        ans = max(ans, f3)
    }

    return int64(ans)
}

        if x < y {
            f2 = max(f2, f1) + x
        } else {
            f2 = math.MinInt / 2
        }

        if x > y {
            f1 = max(f1, y) + x
        } else {
            f1 = math.MinInt / 2
        }

        ans = max(ans, f3)
    }

    return int64(ans)
}