func minLength(s string, numOps int) int {
    n := len(s)

    check := func(mid int) bool {
        cnt := 0
        if mid == 1 {
            for i := range s {
                if (i % 2 == 0) != (s[i] == s[0]) {
                    cnt++
                }
            }
            cnt = min(cnt, n - cnt)
        } else {
            k := 0
            for i := range s {
                k++
                if i == n - 1 || s[i] != s[i + 1] {
                    cnt += k / (mid + 1)
                    k = 0
                }
            }
        }

        return cnt <= numOps
    }
    
    l, r := 0, n
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}