func maxValidSplits(nums []int) int {
    n := len(nums)
    pre := make([]int, n)
    suf := make([]int, n)
    for i := range n {
        j := n - i - 1
        if i == 0 {
            pre[i] = nums[i]
        } else {
            pre[i] = gcd(pre[i - 1], nums[i])
        }

        if j == n - 1 {
            suf[j] = nums[j]
        } else {
            suf[j] = gcd(suf[j + 1], nums[j])
        }
    }
    all_g := suf[0]
    p := sort.Search(n, func(i int) bool {
        return pre[i] == all_g
    })
    q := sort.SearchInts(suf, all_g + 1)
    ans := max(0, q - p - 1)
    for i := range n {
        if i > 0 && pre[i] == pre[i - 1] {
            continue
        }

        L, R := 0, 0
        if i > 0 {
            L = pre[i - 1]
        }
        if i < n - 1 {
            R = suf[i + 1]
        }

        var G int 
        if R > all_g && L % R == 0 {
            G = R
        } else if L > all_g && R % L == 0 {
            G = L
        } else {
            continue
        }

        g := 0
        for j := range n {
            if i == j {
                continue
            }
            g = gcd(nums[j], g)
            if g == G {
                p = j
                break
            }
        }
        g = 0
        for j := n - 1; j >= 0; j-- {
            if i == j {
                continue
            }
            g = gcd(nums[j], g)
            if g == G {
                q = j
                break
            }
        }

        cur := q - p 
        if p <= i && i < q {
            cur--
        }
        ans = max(ans, cur)
        break
    }

    return ans
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }

    return x
}