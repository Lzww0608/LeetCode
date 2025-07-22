func minMutation(startGene string, endGene string, bank []string) int {
    m := make(map[string]bool)
    for _, s := range bank {
        m[s] = true
    }
    if startGene == endGene {
        return 0
    }
    if !m[endGene] {
        return -1
    } 

    check := func(s, t string) bool {
        cnt := 0
        for i := range s {
            if s[i] != t[i] {
                cnt++
            }
        }

        return cnt == 1
    } 

    q := []string{startGene}
    d := 0
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            if cur == endGene {
                return d
            }
            q = q[1:]
            for k, v := range m {
                if v && check(cur, k) {
                    m[k] = false
                    q = append(q, k)
                }
            }
        }   
        d++
    }

    return -1
}