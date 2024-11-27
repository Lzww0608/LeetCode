func maxRepeating(s string, t string) (ans int) {
    n := len(t)
    cnt := 0
    pi := make([]int, n)
    for i := 1; i < n; i++ {
        for cnt > 0 && t[cnt] != t[i] {
            cnt = pi[cnt-1]
        }
        if t[cnt] == t[i] {
            cnt++
        }
        pi[i] = cnt
    }

    pos := []int{}
    cnt = 0
    for i := range s {
        for cnt > 0 && t[cnt] != s[i] {
            cnt = pi[cnt-1]
        }
        if t[cnt] == s[i] {
            cnt++
        }
        if cnt == n {
            pos = append(pos, i - n + 1)
            cnt = pi[cnt-1]
        }
    }

    if len(pos) == 0 {
        return 
    }

    cur := 1
    m := map[int]int{}
    m[pos[0]] = 1
    for i := 1; i < len(pos); i++ {
        cur = m[pos[i]-n] + 1
        m[pos[i]] = cur
        ans = max(ans, cur)
    }
    ans = max(ans, cur)

    return 
}