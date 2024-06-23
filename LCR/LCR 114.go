func alienOrder(words []string) string {
    g := map[byte][]byte{}
    in := map[byte]int{}

    for i := range words[0] {
        c := words[0][i]
        in[c] = 0
    }

    n := len(words)
next:
    for i := 1; i < n; i++ {
        s := words[i-1]
        t := words[i]
        for _, c := range t {
            in[byte(c)] += 0
        }
        for j := 0; j < len(s) && j < len(t); j++ {
            if s[j] != t[j] {
                g[s[j]] = append(g[s[j]], t[j])
                in[t[j]]++
                continue next
            }
        }
        if len(s) > len(t) {
            return ""
        }
    }

    ans := make([]byte, len(in))
    q := ans[:0]
    for u, d := range in {
        if d == 0 {
            q = append(q, u)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            if in[v]--; in[v] == 0 {
                q = append(q, v)
            }
        }
    }

    if cap(q) == 0 {
        return string(ans)
    }

    return ""
}
