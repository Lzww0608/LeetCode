func minOperations(s string) int {
    n := len(s)

    solve := func(l, r int) (res int) {
        for range n / 2 {
            x := int(s[l] - 'a')
            y := int(s[r] - 'a')
            if x < y {
                res += min(y - x, x + 26 - y)
            } else {
                res += min(x - y, y + 26 - x)
            }

            l = (l + 1) % n 
            r = (r - 1 + n) % n
        }

        return
    }

    ans := math.MaxInt 
    for i := range n {
        ans = min(ans, i + solve(i, (i + n - 1) % n))
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}