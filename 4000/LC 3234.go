func numberOfSubstrings(s string) (ans int) {
    n := len(s)
    a := []int{}
    for i, c := range s {
        if c == '0' {
            a = append(a, i)
        }
    }

    one := n - len(a)
    a = append(a, n)
    for i, j := 0, 0; i < n; i++ {
        if s[i] == '1' {
            ans += a[j] - i
        }

        for k := j; k < len(a) - 1; k++ {
            cnt0 := k - j + 1
            if cnt0 * cnt0 > one {
                break
            }

            cnt1 := a[k] - i - (k - j)
            ans += max(a[k + 1] - a[k] - max(cnt0 * cnt0 - cnt1, 0), 0)
        }

        if s[i] == '0' {
            j++
        }
    }

    return
}
