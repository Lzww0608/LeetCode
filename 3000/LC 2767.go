var pow []string

func init() {
    for i := 1; i < (1 << 15); i *= 5 {
        pow = append(pow, strconv.FormatUint(uint64(i), 2))
    }
}

func minimumBeautifulSubstrings(s string) int {
    n := len(s)
    f := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        f[i] = n + 1
        if s[i] == '0' {
            continue
        }
        for _, t := range pow {
            if i + len(t) > n {
                break
            }
            if s[i:i+len(t)] == t {
                f[i] = min(f[i], f[i+len(t)] + 1)
            }
        }
    }

    if f[0] == n + 1 {
        return -1
    }
    return f[0]
}