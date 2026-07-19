func minCost(source string, target string, rules [][]string, costs []int) int {
    n := len(source)
    f := make([]int, n + 1)
    for i := range f {
        f[i] = -1
    }
    f[0] = 0
    for i := range n {
        if f[i] == -1 {
            continue
        }
        if source[i] == target[i] {
            f[i + 1] = f[i]
        }

        for j := range costs {
            pattern, replacement := rules[j][0], rules[j][1]
            m := len(pattern)
            if i + m > n {
                continue
            }

            b := true
            cnt := 0
            for k := range m {
                if (pattern[k] != source[i + k] && pattern[k] != '*') || replacement[k] != target[i + k] {
                    b = false
                    break
                } else if pattern[k] == '*' {
                    cnt++
                }
            }

            if b && (f[i + m] == -1 || f[i + m] > f[i] + cnt + costs[j]) {
                fmt.Println(i, j)
                f[i + m] = f[i] + cnt + costs[j]
            }
        }
    }
    
    return f[n]
}