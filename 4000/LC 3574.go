func maxGCDScore(nums []int, k int) int64 {
    ans := 0
    n := len(nums)

    for i := range n {
        g, cnt := 0, 0
        lb := math.MaxInt
        for j := i; j >= 0; j-- {
            x := nums[j]
            if t := x & (-x); t < lb {
                lb, cnt = t, 1
            } else if t == lb {
                cnt++
            }
            
            g = gcd(x, g)
            if cnt <= k {
                ans = max(ans, g * 2 * (i - j + 1))
            } else {
                ans = max(ans, g * (i - j + 1))
            } 
        }
    }

    return int64(ans)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
}