// a / b == d / c
func numberOfSubsequences(nums []int) int64 {
    type pair struct {
        x, y int
    }

    n := len(nums)
    suf := make(map[pair]int)
    for i := n - 3; i >= 4; i-- {
        x := nums[i]
        for j := i + 2; j < n; j++ {
            y := nums[j]
            g := gcd(x, y)
            suf[pair{y / g, x / g}]++
        }
    }

    ans := 0
    for i := 2; i < n - 4; i++ {
        x := nums[i]
        for j := i - 2; j >= 0; j-- {
            y := nums[j]
            g := gcd(x, y)
            
            ans += suf[pair{y / g, x / g}]
        }

        x = nums[i + 2]
        for _, y := range nums[i + 4:] {
            g := gcd(x, y)
            suf[pair{y / g, x / g}]--
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