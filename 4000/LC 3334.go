func maxScore(nums []int) int64 {
    n := len(nums)
    suf := make([]int, n + 1)
    sufLcm := make([]int, n + 1)
    sufLcm[n] = 1
    for i := 0; i < n; i++ {
        suf[n-i-1] = gcd(nums[n-i-1], suf[n-i])
        sufLcm[n-i-1] = lcm(nums[n-i-1], sufLcm[n-i])
    }
    ans := suf[0] * sufLcm[0]
    pre, preLcm := 0, 1
    for i, x := range nums {
        ans = max(ans, gcd(pre, suf[i+1]) * lcm(preLcm, sufLcm[i+1]))
        pre = gcd(x, pre)
        preLcm = lcm(preLcm, x)
    }
    return int64(ans)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y 
    }
    return x 
}

func lcm(x, y int) int {
    return x * y / gcd(x, y)
}