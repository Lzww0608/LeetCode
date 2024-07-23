const N int = 10_000
func openLock(deadends []string, target string) int {
    m := map[string]bool{}
    for _, s := range deadends {
        m[s] = true
    }

    if m[target] || m["0000"] {
        return -1
    } else if target == "0000" {
        return 0
    }

    q := []string{"0000"}
    vis := map[string]bool{"0000": true}
    d := 1
    for len(q) > 0 {
        for t := len(q); t > 0; t-- {
            cur := []byte(q[0])
            q = q[1:]
            for i := range cur {
                tmp := int(cur[i] - '0')
                for x := -1; x <= 1; x += 2 {
                    cur[i] = byte('0' + (tmp + x + 10) % 10) 
                    s := string(cur)
                    if s == target {
                        return d
                    }
                    if !vis[s] && !m[s] {
                        vis[s] = true
                        q = append(q, s)
                    }

                }
                cur[i] = byte(tmp + '0')
            }
        }
        d++
    }

    return -1
}