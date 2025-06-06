const N int = 26
const M int = 8
func longestSubsequenceRepeatedK(s string, k int) string {
    n := len(s)
    pos := [N]int{}
    for i := range N {
        pos[i] = n
    }

    cnt := [N]int{}
    nxt := make([][N]int, n)
    for i := n - 1; i >= 0; i-- {
        nxt[i] = pos
        x := int(s[i] & 31) - 1
        pos[x] = i 
        cnt[x]++
    }

    a := []byte{}
    for i := N - 1; i >= 0; i-- {
        if cnt[i] >= k {
            a = append(a, byte('a' + i))
        }
    }

    check := func(t string) bool {
        i, j := 0, 0
        if s[0] == t[0] {
            j = 1
        }

        for j < k * len(t) {
            i = nxt[i][t[j % len(t)] - 'a']
            if i == n {
                return false
            }
            j++
        }

        return true
    }

    ans := [M][]string{}
    ans[0] = append(ans[0], "")
    for t := 1; t < M; t++ {
        for _, p := range ans[t-1] {
            for _, c := range a {
                tmp := p + string(c)
                if check(tmp) {
                    ans[t] = append(ans[t], tmp)
                }
            }
        }
    }

    for t := M - 1; t > 0; t-- {
        if len(ans[t]) > 0 {
            return ans[t][0]
        }
    }

    return ""
}