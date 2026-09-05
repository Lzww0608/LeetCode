func maxValidSplits(nums []int) (ans int) {
    n := len(nums)
    for i := range n + 1 {
        cur := 0
        suf := make([]int, n + 1)
        for j := n - 1; j >= 0; j-- {
            if j == i {
                suf[j] = suf[j + 1] 
            } else {
                suf[j] = gcd(nums[j], suf[j + 1])
            }
        }

        pre := 0
        for j := range n {
            if j == i {
                continue
            }
            pre = gcd(nums[j], pre)
            if pre == suf[j + 1] {
                cur++
            }
        }

        ans = max(ans, cur)
    }

    return
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}