const N int = 26
func smallestPalindrome(s string, k int) string {
    n := len(s)
    cnt := [N]int{}
    for i := range n / 2 {
        x := int(s[i] - 'a')
        cnt[x]++
    }

    comb := func(n, m int) int {
        m = min(m, n - m)
        res := 1
        for i := 1; i <= m; i++ {
            res = res * (n - m + i) / i
            if res >= k {
                return res
            }
        }

        return res
    } 

    perm := func(m int) int {
        res := 1
        for _, x := range cnt {
            res *= comb(m, x)
            if res >= k {
                return res
            }
            m -= x
        }

        return res
    }

    if perm(n / 2) < k {
        return ""
    }

    ans := make([]byte, n)
    ans[n / 2] = s[n / 2]
    for i := range n / 2 {
        for j := range N {
            if cnt[j] == 0 {
                continue
            }
            cnt[j]--
            p := perm(n / 2 - i - 1)
            if p >= k {
                ans[i], ans[n - i - 1] = byte('a' + j), byte('a' + j)
                break
            }
            k -= p 
            cnt[j]++
        }
    }

    return string(ans)
}