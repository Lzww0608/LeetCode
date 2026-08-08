func maxPairStrength(nums []int) int64 {
    n := len(nums)
    ans := 0
    for i := range nums {
        for j := i + 1; j < n; j++ {
            t := gcd(nums[i], nums[j])
            ans = max(ans, (nums[i] * nums[j]) / (t * t))
        }
    }

    return int64(ans)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }

    return x
}