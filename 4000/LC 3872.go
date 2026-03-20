func longestArithmetic(a []int) int {
    ans := 1 
    n := len(a)
    suf := make([]pair, n)
    suf[n - 1].l = 1
    for i := n - 2; i >= 0; i-- {
        x := a[i + 1] - a[i]
        suf[i].d = x
        if x == suf[i + 1].d {
            suf[i].l = suf[i + 1].l + 1
        } else {
            suf[i].l = 2
        }

        if i - 2 >= 0 && a[i] - a[i - 2] == x * 2 {
            ans = max(ans, 2 + suf[i].l)
        }
    }

    d, l := 0, 1
    for i := 1; i < n; i++ {
        x := a[i] - a[i - 1]
        if x == d {
            l++
        } else {
            d, l = x, 2
        }
        ans = max(ans, l + 1)

        if i + 2 < n && a[i + 2] - a[i] == d * 2 {
            ans = max(ans, l + 2)
            if d == suf[i + 2].d {
                ans = max(ans, l + 1 + suf[i + 2].l)
            }
        }
    }

    return min(ans, n)
}

type pair struct {
    d, l int 
}